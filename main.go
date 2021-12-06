package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>Email me at <a href=\"mailto:example@gmail.com\">example@gmail.com</a></p>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.NotFound(w, r)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)
	}

	fmt.Printf("Path is: %s\n", r.URL.Path)
}

func main() {

	// // manually setting up routing for each endpoint
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)

	// // customer routing using the HandlerFunc type, which implements the Handler interface
	// // var router http.HandlerFunc = pathHandler
	// router := http.HandlerFunc(pathHandler)

	// custom routing using the Handler interface
	router := Router{}

	fmt.Println("Starting the server on :8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
