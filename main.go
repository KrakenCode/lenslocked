package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// make the path OS agnostic
	tplPath := filepath.Join("templates", "home.gohtml")

	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("error parsing home template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("error executing home template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>Email me at <a href=\"mailto:example@gmail.com\">example@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := `
	<h1>FAQ Page</h1>
	<p>
	Q: Is there a free version?<br>
	A: Yes! We offer a free trial for 30 days on any paid plans.<br><br>

	Q: What are your support hours?<br>
	A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.<br><br>

	Q: How do I contact support?<br>
	A: Email us = support@example.com<br><br>
	</p>
	`

	fmt.Fprint(w, content)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	content := fmt.Sprintf("<p>Passed in resource id is: %s</p>", id)
	fmt.Fprint(w, content)
}
