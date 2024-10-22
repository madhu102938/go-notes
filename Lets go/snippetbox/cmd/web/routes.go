package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// creating a new server mux which means a new router
	mux := http.NewServeMux()

	// subtree path
	mux.HandleFunc("/", app.MainHandler)
	// fixed path
	mux.HandleFunc("/snippet", app.SnippetHandler)
	// fixed path
	mux.HandleFunc("/snippet/create", app.CreateHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/download/", app.DownloadHandler)

	return mux
}