package main

import (
	"net/http"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// creating a new server mux which means a new router
	mux := http.NewServeMux()

	// subtree path (no wildcard here)
	mux.HandleFunc("/", app.MainHandler) // exception, only works for '/' and nothing else

	mux.HandleFunc("/snippet/create", app.CreateHandler)
	mux.HandleFunc("/snippet", app.SnippetHandler)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	myChain := alice.New(app.logRequest, secureHeaders)
	myOtherChain := myChain.Append(app.recoverPanic)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return myOtherChain.Then(mux)
}