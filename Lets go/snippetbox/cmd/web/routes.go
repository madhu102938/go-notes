package main

import (
	"net/http"
	"github.com/justinas/alice"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	// creating a new server mux which means a new router
	mux := pat.New()

	// subtree path (no wildcard here)
	mux.Get("/", http.HandlerFunc(app.MainHandler)) // exception, only works for '/' and nothing else
	// mux.Post("/snippet/create", http.HandlerFunc(app.CreateSnippetForm))
	mux.Get("/snippet/create", http.HandlerFunc(app.CreateHandler))
	mux.Get("/snippet/:id", http.HandlerFunc(app.SnippetHandler))	// `:id` Moved down

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	myChain := alice.New(app.logRequest, secureHeaders)
	myOtherChain := myChain.Append(app.recoverPanic)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return myOtherChain.Then(mux)
}