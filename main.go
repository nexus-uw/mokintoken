package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Note struct {
	id            string
	encryptedText string
	expiry        string
	expiresViews  int
	viewCount     int
}

func lookUpNote(id string) (Note, error) {

}

func saveNote(encryptedText string, expiry string, expiresViews int) string {

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

func decryptHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	// db lookup id
	note, err := lookUpNote(id)

	if err != nil {
		// todo use nice template
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

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

func saveNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "/api/save-note Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")

	http.HandleFunc("/", ass("home"))
	http.HandleFunc("/about", ass("about"))
	http.HandleFunc("/decrypt/", decryptHandler)
	http.HandleFunc("/api/save-note", saveNoteHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
