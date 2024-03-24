package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Fileserver to serve static files (HTML, CSS, JS)
	fileServer := http.FileServer(http.Dir("client/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// URL Routing
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.snippetCreate)
	mux.HandleFunc("/snippet/view", app.snippetView)

	return mux
}
