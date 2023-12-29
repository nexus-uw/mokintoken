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
	expiry        string
	expiresViews  int
	viewCount     int
}

func lookUpNote(db *sql.DB, id string) Note {
	row := db.QueryRow("SELECT encryptedText, expiry, expiresViews, viewCount  FROM notes WHERE id = ? LIMIT 1", id)
	// todo cleanly account for NOT FOUND
	var (
		encryptedText string
		expiry        string
		expiresViews  int
		viewCount     int
	)
	err := row.Scan(&encryptedText, &expiry, &expiresViews, &viewCount)

	if err != nil {
		panic(err)
	}

	return Note{id, encryptedText, expiry, expiresViews, viewCount}

}

func incrementViews(db *sql.DB, id string) {
	_, err := db.Exec(" TODO id = ?", id)

	if err != nil {
		panic(err)
	}
}

func deleteNote(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM notes WHERE id = ?", id)

	if err != nil {
		panic(err)
	}
}

func saveNote(db *sql.DB, encryptedText string, expiry string, expiresViews int) string {
	var id = uuid.New().String()

	_, err := db.Exec("INSERT INTO Notes (id, encryptedText, expiry, expiresViews, viewCount) VALUES (?,?, ?,?,?)", id, encryptedText, expiry, expiresViews, 0)

	if err != nil {
		panic(err)
	}

	return id
}

func ass(name string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Use the template.ParseFiles() function to read the files and store the
		// templates in a template set. Notice that we use ... to pass the contents
		// of the files slice as variadic arguments.
		ts, err := template.ParseFiles(
			"./views/template.html",
			"./views/"+name+".html",
		)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Use the ExecuteTemplate() method to write the content of the "base"
		// template as the response body.
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func decryptHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ff := strings.Split(r.URL.Path, "/")
		id := ff[len(ff)-1]
		// db lookup id
		note /*, err*/ := lookUpNote(db, id)

		/*if err != nil {
			// todo use nice template
			http.Error(w, "note not found", http.StatusNotFound)
			return
		}*/
		log.Println(note)
		if note.expiresViews <= note.viewCount || note.expiry <= fmt.Sprint(time.Now().UTC().UnixMilli()) {
			http.Error(w, "note not found", http.StatusNotFound)
			return // todo nice template
		}
		// todo: inspect view count + expiry time
		// if expired, delete note + return 404
		// else inc view count
		incrementViews(db, note.id)

		// Use the template.ParseFiles() function to read the files and store the
		// templates in a template set. Notice that we use ... to pass the contents
		// of the files slice as variadic arguments.
		ts, err := template.ParseFiles(
			"./views/template.html",
			"./views/decrypt.html",
		)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Use the ExecuteTemplate() method to write the content of the "base"
		// template as the response body.
		err = ts.ExecuteTemplate(w, "base", note.encryptedText)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

type NewNote struct {
	encryptedText string `json:"encryptedText"`
	expiryString  string `json:"expiryString"`
	expressViews  string `json:"expressViews"`
}

func saveNoteHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		
		var t NewNote
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		views, _ := strconv.Atoi(t.expressViews)
		log.Println(t)
		id := saveNote(db, t.encryptedText, t.expiryString, views)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"id\":\"%s\"}", id) // todo: do real encoding + proper json

	}
}

func setUpDB(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Notes
	(id TEXT PRIMARY KEY,
	encryptedText TEXT NOT NULL,
	expiry TEXT NOT NULL,
	expiresViews INT DEFAULT 2,
	viewCount INT DEFAULT 0)`)
	// todo: come up with migration process
	if err != nil {
		panic(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}
	log.Println(" setting up db")
	setUpDB(db)
	log.Println(" db set up")

	http.HandleFunc("/", ass("home"))
	http.HandleFunc("/about", ass("about"))
	http.HandleFunc("/decrypt/", decryptHandler(db))
	http.HandleFunc("/api/save-note", saveNoteHandler(db))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
