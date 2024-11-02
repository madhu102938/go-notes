package main

import (
	"net/http"
	"github.com/justinas/alice"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	// creating a new server mux which means a new router
	mux := pat.New()

	dynamicMiddleware := alice.New(app.session.Enable)

	// subtree path (no wildcard here)
	mux.Get("/", dynamicMiddleware.ThenFunc(app.MainHandler)) // exception, only works for '/' and nothing else
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.CreateHandler))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.CreateSnippetForm))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.SnippetHandler))	// `:id` Moved down

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	myChain := alice.New(app.logRequest, secureHeaders)
	myOtherChain := myChain.Append(app.recoverPanic)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return myOtherChain.Then(mux)
}