// Starting from the book "Lets go"

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// creating a new server mux which means a new router
	mux := http.NewServeMux()

	// subtree path
	mux.HandleFunc("/", MainHandler)
	// fixed path
	mux.HandleFunc("/snippet", SnippetHandler)
	// fixed path
	mux.HandleFunc("/snippet/create", CreateHandler)
	
	fmt.Println("Server started on 8080")
	
	// by default, localhost is used
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("create a new snippet"))
}

func SnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a snippet"))
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// if the path is not the root path, then return a 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Display the home page"))
}

