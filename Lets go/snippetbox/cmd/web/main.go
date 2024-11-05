// Starting from the book "Lets go"

package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"crypto/tls"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"

	"snippetbox/pkg/models/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
 type application struct {
	infoLog			*log.Logger
	errorLog		*log.Logger
	snippets		*mysql.SnippetModel
	templateCache	map[string]*template.Template
	session			*sessions.Session
	users			*mysql.UserModel
 }

func main() {
	// Errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Commandline arguments
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "web:snipass@/snippetbox?parseTime=true", "MySQL data source name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret Key")

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

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	
	app := &application{
		infoLog:		infoLog,
		errorLog: 		errorLog,
		snippets: 		&mysql.SnippetModel{DB:db},
		templateCache: 	templateCache,
		session: 		session,
		users:			&mysql.UserModel{DB: db},
	}

	tlsConflig := &tls.Config{
		PreferServerCipherSuites: 	true,
		CurvePreferences:			[]tls.CurveID{tls.X25519, tls.CurveP256},
	}
	
	srv := &http.Server{
		Addr:*addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
		TLSConfig: tlsConflig,
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	// `addr` is pointer to string, thus we need to dereference it
	infoLog.Println("Server started on", *addr)
	
	// by default, localhost is used
	// err = srv.ListenAndServe()
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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

