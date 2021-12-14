package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KrakenCode/lenslocked/controllers"
	"github.com/KrakenCode/lenslocked/templates"
	"github.com/KrakenCode/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tpl := views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "signup.gohtml"))
	r.Get("/signup", controllers.StaticHandler(tpl))

	fmt.Println("Starting the server on localhost:8080...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Panicln(err)
	}
}
