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

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/download/", DownloadHandler)
	
	fmt.Println("Server started on 8080")
	
	// by default, localhost is used
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

