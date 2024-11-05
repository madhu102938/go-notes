package main

import (
	"net/http"
	"github.com/justinas/alice"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {
	// creating a new server mux which means a new router
	mux := pat.New()

	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	// subtree path (no wildcard here)
	mux.Get("/", dynamicMiddleware.ThenFunc(app.MainHandler)) // exception, only works for '/' and nothing else
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.CreateHandler))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.CreateSnippetForm))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.SnippetHandler))	// `:id` Moved down

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserFrom))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	myChain := alice.New(app.logRequest, secureHeaders)
	myOtherChain := myChain.Append(app.recoverPanic)

	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return myOtherChain.Then(mux)
}