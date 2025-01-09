package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Program 2: Simple REST API (Setup as pre-written code)
// Install gorilla/mux: go get -u github.com/gorilla/mux

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go REST API!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/welcome", homeHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
