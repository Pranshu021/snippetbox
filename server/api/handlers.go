package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Calling centralized error handlers
		app.notFound(w)
		return
	}

	template_files := []string{
		"client/ui/html/base.tmpl",
		"client/ui/html/partials/nav.tmpl",
		"client/ui/html/home.tmpl",
	}

	ts, err := template.ParseFiles(template_files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// app.errorLog.Printf(err.Error())
		app.serverError(w, err)
		return
	}

	w.Write([]byte("Welcome to Homepage"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// Sending client error with centralzied error handlers
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a snippet"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display the snippet of id %d", id)
}
