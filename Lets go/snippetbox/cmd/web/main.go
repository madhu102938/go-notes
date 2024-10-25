// Starting from the book "Lets go"

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox/pkg/models/mysql"
)

 // Define an application struct to hold the application-wide dependencies for the
 // web application. For now we'll only include fields for the two custom loggers, but
 // we'll add more to it as the build progresses.
 type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
	snippets *mysql.SnippetModel
 }

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:snipass@/snippetbox?parseTime=true", "MySQL data source name")

	//  You need to call this *before* you use the addr variable
    // otherwise it will always contain the default value of ":4000". If any errors are
    // encountered during parsing the application will be terminated
	flag.Parse()
	
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err.Error())
	}
	
	defer db.Close()

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
		snippets: &mysql.SnippetModel{DB:db},
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

