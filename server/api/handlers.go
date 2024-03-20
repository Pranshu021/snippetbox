package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	template_files := []string{
		"client/ui/html/base.tmpl",
		"client/ui/html/partials/nav.tmpl",
		"client/ui/html/home.tmpl",
	}

	ts, err := template.ParseFiles(template_files...)
	if err != nil {
		app.errorLog.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errorLog.Printf(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Write([]byte("Welcome to Homepage"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Mothod Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a snippet"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display the snippet of id %d", id)

	// w.Write([]byte("View snippet"))
}
