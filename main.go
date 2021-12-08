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

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	fmt.Println("Starting the server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panicln(err)
	}
}
