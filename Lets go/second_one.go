// Starting from the book "Lets go"

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

var count int = 0
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		count++;
	}()

	if r.Method != http.MethodPost {
		if count == 0 {
			w.Header().Set("allow-hehe", http.MethodPost)
		}
		if count == 1 {
			w.Header().Add("Allow", "hehe")
		}
		// fmt.Println(w.Header().Get("Allow"))
		if count > 2 {
			w.Header().Del("Allow")
		}
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "method not allowed", 405)
		return
	}

	w.Write([]byte("create a new snippet"))
}

func SnippetHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Display a snippet"))
	// receiving arguements
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		fmt.Fprintln(w, "not a valid id")
		return
	}

	fmt.Fprintf(w, "Displaying a specific snippet with ID %v\n", id)
	// writing to response writer with fprintf
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// if the path is not the root path, then return a 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Display the home page"))
}

