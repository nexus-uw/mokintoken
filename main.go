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
}

func lookUpNote(db *sql.DB, id string) (Note, error) {
	row := db.QueryRow(`SELECT encryptedText, expiry, expiresViews, viewCount  
	FROM notes 
	WHERE id = ? 
	LIMIT 1`, id)
	// todo cleanly account for NOT FOUND
	var (
		encryptedText string
		expiry        time.Time
		expiresViews  int
		viewCount     int
	)
	err := row.Scan(&encryptedText, &expiry, &expiresViews, &viewCount)

	if err == sql.ErrNoRows {
		return Note{}, err
	}
	if err != nil {
		panic(err)
	}

	return Note{id, encryptedText, expiry, expiresViews, viewCount}, nil

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

func saveNote(db *sql.DB, encryptedText string, expiry time.Time, expiresViews int) string {
	var id = uuid.New().String()

	_, err := db.Exec("INSERT INTO Notes (id, encryptedText, expiry, expiresViews, viewCount) VALUES (?,?, ?,?,?)", id, encryptedText, expiry, expiresViews, 0)

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



		err = ts.ExecuteTemplate(w, "base", DecryptTemplateContext{os.Getenv("CLEARNET"), os.Getenv("DARKNET"), note.encryptedText})
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
}

func saveNoteHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		var t NewNote
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		views, _ := strconv.Atoi(t.ExpiresViews)
		expiryHours, _ := strconv.Atoi(t.ExpiresHours)
		expiry := time.Now().Add(time.Hour * time.Duration(expiryHours))
		id := saveNote(db, t.EncryptedText, expiry, views)
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Notes
	(id TEXT PRIMARY KEY,
	encryptedText TEXT NOT NULL,
	expiry DATETIME NOT NULL,
	expiresViews INT DEFAULT 1,
	viewCount INT DEFAULT 0)`)
	// todo: come up with migration process
	if err != nil {
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
	http.HandleFunc("/decrypt/", decryptHandler(db))
	http.HandleFunc("/api/save-note", saveNoteHandler(db))
	http.HandleFunc("/ping", ping(db))
	// this should be handled by a cdn
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
