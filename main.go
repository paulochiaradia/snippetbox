package main

import (
	"log"
	"net/http"
)

// homeHandler writes a byte slice containing "Hello from Snppetbox"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//Check if the current request URL path exactly matches "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// snippetView Display a specifc snippet
func snippetViewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specifc snippet"))
}

// snippetCreate Create a new snippet
func snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)
	log.Println("Stating server on:4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
