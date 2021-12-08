package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/KrakenCode/lenslocked/controllers"
	"github.com/KrakenCode/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Panicln(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Panicln(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Panicln(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	fmt.Println("Starting the server on :8080...")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panicln(err)
	}
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
