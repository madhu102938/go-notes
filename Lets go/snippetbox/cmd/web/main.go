// Starting from the book "Lets go"

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	addr := flag.String("addr", ":8080", "HTTP network address")

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

	srv := &http.Server{
		Addr:*addr,
		Handler: mux,
		ErrorLog: errorLog,
	}
	
	// `addr` is pointer to string, thus we need to dereference it
	infoLog.Println("Server started on", *addr)
	
	// by default, localhost is used
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

