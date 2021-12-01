package main

import (
	"fmt"
	"net/http"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", homeHandlerFunc)

	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
