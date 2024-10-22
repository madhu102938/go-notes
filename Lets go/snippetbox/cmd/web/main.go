// Starting from the book "Lets go"

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	/* we can also use
	var addr *string = new(string)  // new(Type) returns pointer to Type, allocates memory
	flag.StringVar(addr, "addr", ":8080", "HTTP network address")
	*/

	//  You need to call this *before* you use the addr variable
    // otherwise it will always contain the default value of ":4000". If any errors are
    // encountered during parsing the application will be terminated
	flag.Parse()

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
	
	// `addr` is pointer to string, thus we need to dereference it
	log.Println("Server started on", *addr)
	
	// by default, localhost is used
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

