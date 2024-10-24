// Starting from the book "Lets go"

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

 // Define an application struct to hold the application-wide dependencies for the
 // web application. For now we'll only include fields for the two custom loggers, but
 // we'll add more to it as the build progresses.
 type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
 }

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
	}


	addr := flag.String("addr", ":8080", "HTTP network address")

	//  You need to call this *before* you use the addr variable
    // otherwise it will always contain the default value of ":4000". If any errors are
    // encountered during parsing the application will be terminated
	flag.Parse()

	srv := &http.Server{
		Addr:*addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
	}
	
	// `addr` is pointer to string, thus we need to dereference it
	infoLog.Println("Server started on", *addr)
	
	// by default, localhost is used
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

