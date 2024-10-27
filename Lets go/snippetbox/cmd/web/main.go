// Starting from the book "Lets go"

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"database/sql"
	"html/template"
	
	"snippetbox/pkg/models/mysql"
	
	_ "github.com/go-sql-driver/mysql"
)

 // Define an application struct to hold the application-wide dependencies for the
 // web application. For now we'll only include fields for the two custom loggers, but
 // we'll add more to it as the build progresses.
 type application struct {
	infoLog			*log.Logger
	errorLog		*log.Logger
	snippets		*mysql.SnippetModel
	templateCache	map[string]*template.Template
 }

func main() {
	// Errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Commandline arguments
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:snipass@/snippetbox?parseTime=true", "MySQL data source name")

	//  You need to call this *before* you use the addr variable
    // otherwise it will always contain the default value of ":4000". If any errors are
    // encountered during parsing the application will be terminated
	flag.Parse()
	
	// initiating mysql connection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err.Error())
	}
	
	defer db.Close()

	// caching
	templateCache, err := newTemplateCache("./ui/html") // path relative to home of package
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
		snippets: &mysql.SnippetModel{DB:db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:*addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
	}
	
	// `addr` is pointer to string, thus we need to dereference it
	infoLog.Println("Server started on", *addr)
	
	// by default, localhost is used
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// db.SetMaxIdleConns(X)
	// db.SetMaxOpenConns(Y)

	return db, nil
}
