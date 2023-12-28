package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
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

func lookUpNote(db *sql.DB, id string) (Note) {
	row := db.QueryRow("SELECT encryptedText, expiry, expiresViews, viewCount  FROM notes WHERE id = ? LIMIT 1", id)
	// todo cleanly account for NOT FOUND
	var (
	encryptedText string
	expiry        string
	expiresViews  int
	viewCount     int
	)
	err := row.Scan(&encryptedText, &expiry, &expiresViews, &viewCount )

	if err != nil{
		panic(err)
	}

	return Note{id, encryptedText, expiry, expiresViews, viewCount}
	
}

func incrementViews(db *sql.DB, id string)  {
	_, err:= db.Exec(" TODO id = ?", id)
	
	if err != nil{
		panic(err)
	}
}

func deleteNote(db *sql.DB, id string)  {
	_, err:= db.Exec("DELETE FROM notes WHERE id = ?", id)
	
	if err != nil{
		panic(err)
	}
}

func saveNote(db *sql.DB,encryptedText string, expiry string, expiresViews int) string {
	var id = uuid.New().String()

	_, err:= db.Exec("INSERT INTO notes (id, encryptedText, expiry, expiresViews, viewCount) VALUES (?,?, ?,?,?)", id,encryptedText, expiry, expiresViews, 0)
	
	if err != nil{
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

func decryptHandler(db *sql.DB) func (w http.ResponseWriter, r *http.Request){
	return func (w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[1:]
		// db lookup id
		note/*, err*/ := lookUpNote(db, id)

		/*if err != nil {
			// todo use nice template
			http.Error(w, "note not found", http.StatusNotFound)
			return
		}*/
		if (note.expiresViews <= note.viewCount || note.expiry <= fmt.Sprint(time.Now().UTC().UnixMilli()) ){
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

func saveNoteHandler(db *sql.DB) func (w http.ResponseWriter, r *http.Request) {
	return  func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "/api/save-note Hi there, I love %s!", r.URL.Path[1:])
	 // saveNote
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		panic(err)
	}

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
