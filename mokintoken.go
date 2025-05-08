package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	id            string
	encryptedText string
	expiry        time.Time
	expiresViews  int
	viewCount     int
	encryptedImg  string
}

func lookUpNote(db *sql.DB, id string) (Note, error) {
	row := db.QueryRow(`SELECT encryptedText, expiry, expiresViews, viewCount, encryptedImg
	FROM notes
	WHERE id = ?
	LIMIT 1`, id)
	// todo cleanly account for NOT FOUND
	var (
		encryptedText string
		expiry        time.Time
		expiresViews  int
		viewCount     int
		encryptedImg  string
	)
	err := row.Scan(&encryptedText, &expiry, &expiresViews, &viewCount, &encryptedImg)

	if err == sql.ErrNoRows {
		return Note{}, err
	}
	if err != nil {
		panic(err)
	}

	return Note{id, encryptedText, expiry, expiresViews, viewCount, encryptedImg}, nil

}

func incrementViews(db *sql.DB, id string, views int) {
	_, err := db.Exec("UPDATE Notes SET viewCount=? WHERE id=?", views+1, id)

	if err != nil {
		panic(err)
	}
}

func deleteNote(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM Notes WHERE id = ?", id)

	if err != nil {
		panic(err)
	}
}

func saveNote(db *sql.DB, encryptedText string, expiry time.Time, expiresViews int, encryptedImg string) string {
	var id = uuid.New().String()

	_, err := db.Exec("INSERT INTO Notes (id, encryptedText, expiry, expiresViews, viewCount, encryptedImg) VALUES (?,?,?,?,?,?)", id, encryptedText, expiry, expiresViews, 0, encryptedImg)

	if err != nil {
		panic(err)
	}

	return id
}

func userViews(name string) func(w http.ResponseWriter, r *http.Request) {
	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we use ... to pass the contents
	// of the files slice as variadic arguments.
	ts, err := template.ParseFiles(
		"./views/template.html",
		"./views/"+name+".html",
	)
	return func(w http.ResponseWriter, r *http.Request) {

		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		type TemplateContext struct {
			Clearnet string
			Darknet  string
		}
		err = ts.ExecuteTemplate(w, "base", TemplateContext{os.Getenv("CLEARNET"), os.Getenv("DARKNET")})
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

type DecryptTemplateContext struct {
	Clearnet      string
	Darknet       string
	EncryptedText string
	EncryptedImg  string
}

func decryptHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	ts, _ := template.ParseFiles(
		"./views/template.html",
		"./views/decrypt.html",
	)
	ts1, _ := template.ParseFiles(
		"./views/template.html",
		"./views/noteDoesNotExist.html",
	)
	return func(w http.ResponseWriter, r *http.Request) {

		// https://github.com/signalapp/Signal-Android/issues/9958
		if strings.Contains(strings.ToLower(r.Header.Get("User-Agent")), "whatsapp") {
			log.Print("403 - whatsapp/signal")
			w.WriteHeader(http.StatusForbidden)
			ts1.ExecuteTemplate(w, "base", DecryptTemplateContext{})
			return
		}

		ff := strings.Split(r.URL.Path, "/")
		id := ff[len(ff)-1]
		note, err := lookUpNote(db, id)

		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			ts1.ExecuteTemplate(w, "base", DecryptTemplateContext{})
			return
		}
		if note.expiresViews <= note.viewCount || note.expiry.Before(time.Now()) {
			deleteNote(db, note.id)
			w.WriteHeader(http.StatusNotFound)
			ts1.ExecuteTemplate(w, "base", nil)
			return
		}

		incrementViews(db, note.id, note.viewCount)

		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = ts.ExecuteTemplate(w, "base", DecryptTemplateContext{os.Getenv("CLEARNET"), os.Getenv("DARKNET"), note.encryptedText, note.encryptedImg})
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

type NewNote struct {
	EncryptedText string `json:"encryptedText"`
	ExpiresHours  string `json:"expiresHours"`
	ExpiresViews  string `json:"expiresViews"`
	EncryptedImg  string `json:"encryptedImg"`
}

func saveNoteHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		var t NewNote
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		var MAX_FILE_SIZE = 5000000
		if len(t.EncryptedImg) > MAX_FILE_SIZE /*max img size - could be a config param*/ {
			http.Error(w, "img too big (after encryption)", http.StatusBadRequest)
		}
		views, _ := strconv.Atoi(t.ExpiresViews)
		expiryHours, _ := strconv.Atoi(t.ExpiresHours) // should set a max age on this
		expiry := time.Now().Add(time.Hour * time.Duration(expiryHours))
		id := saveNote(db, t.EncryptedText, expiry, views, t.EncryptedImg)
		w.Header().Set("Content-Type", "application/json")

		fmt.Fprintf(w, "{\"id\":\"%s\"}", id) // todo: do real encoding + proper json

	}
}

func ping(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "pong")
	}
}

func setUpDB(db *sql.DB) {
	var err error
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Notes
	(id TEXT PRIMARY KEY,
	encryptedText TEXT NOT NULL,
	expiry DATETIME NOT NULL,
	expiresViews INT DEFAULT 1,
	viewCount INT DEFAULT 0)`)
	// todo: come up with migration process/ lol whats that?
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`ALTER TABLE Notes ADD COLUMN encryptedImg;`)
	if err != nil && err.Error() != `duplicate column name: encryptedImg` {
		panic(err)
	}
}

func cleanUpDb(db *sql.DB) {
	// in place of a good cron job, run clean up at server boot
	_, err := db.Exec(`DELETE FROM Notes WHERE expiry < date('now') OR viewCount >= expiresViews`)
	if err != nil {
		panic(err)
	}
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filename)
    })
}

func main() {

	db, err := sql.Open("sqlite3", "./database/mokintoken.sqlite")

	if err != nil {
		panic(err)
	}
	setUpDB(db)
	log.Println("db set up")
	cleanUpDb(db)
	log.Println("db has been cleaned up of expired notes")

	http.HandleFunc("/", userViews("home"))
	http.HandleFunc("/about", userViews("about"))
	http.HandleFunc("/noteSaved", userViews("noteSaved"))
	http.HandleFunc("/decrypt/", decryptHandler(db))
	http.HandleFunc("/api/save-note", saveNoteHandler(db))
	http.HandleFunc("/ping", ping(db))
	// this should be handled by a cdn
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

  serveSingle("/service-worker.js", "./assets/service-worker.js")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
