package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	//Extract the value of id parameter from the query string and
	//convert to integer. If its not a integer show the not found page
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || ID <= 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", ID)
}

// snippetCreate Create a new snippet
func snippetCreateHandler(w http.ResponseWriter, r *http.Request) {
	//Using r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
