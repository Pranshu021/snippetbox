package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Fileserver to serve static files (HTML, CSS, JS)
	fileServer := http.FileServer(http.Dir("client/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// URL Routing
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/snippet/view", app.snippetView)

	// Using alice library to create chain of middlewares
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Return/Execute standard middleware chain followed by serveMux
	return standard.Then(mux)

	// Wrapping the middlewares as chains
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
