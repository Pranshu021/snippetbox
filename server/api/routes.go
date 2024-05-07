package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Initialize the httpRouter
	router := httprouter.New()

	// Fileserver to serve static files (HTML, CSS, JS)
	fileServer := http.FileServer(http.Dir("client/ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	// URL Routing with httpRouter
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	// Using alice library to create chain of middlewares
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Return/Execute standard middleware chain followed by router
	return standard.Then(router)

	// Wrapping the middlewares as chains
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
