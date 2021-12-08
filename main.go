package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/KrakenCode/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/resource/{id}", resourceHandler)

	fmt.Println("Starting the server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath, nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath, nil)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := fmt.Sprintf("<p>Passed in resource id is: %s</p>", id)
	fmt.Fprint(w, content)
}

func executeTemplate(w http.ResponseWriter, filepath string, data interface{}) {
	tpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("error parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, data)
}
