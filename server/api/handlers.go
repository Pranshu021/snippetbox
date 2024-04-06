package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"snippetbox.sting.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// Calling centralized error handlers
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// template_files := []string{
	// 	"client/ui/html/base.tmpl",
	// 	"client/ui/html/partials/nav.tmpl",
	// 	"client/ui/html/home.tmpl",
	// }

	// ts, err := template.ParseFiles(template_files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	// app.errorLog.Printf(err.Error())
	// 	app.serverError(w, err)
	// 	return
	// }

	// w.Write([]byte("Welcome to Homepage"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"client/ui/html/base.tmpl",
		"client/ui/html/partials/nav.tmpl",
		"client/ui/html/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{
		Snippet: snippet,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}
