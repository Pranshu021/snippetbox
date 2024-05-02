package main

import (
	"html/template"
	"path/filepath"
	"time"

	"snippetbox.sting.net/internal/models"
)

// templateData type to act as a Holding structure for dynamic data to be passed onto templates.
// Can add multiple structs and wrap it into one
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// Custom template function
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object to register our custom template function
var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// A map acting as a cache
	cache := map[string]*template.Template{}

	// Get a slice of filepaths that match the *tmpl pattern in the directory using Glob method
	pages, err := filepath.Glob("client/ui/HTML/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Loop over the filepaths
	for _, page := range pages {
		// Extract the filename from the filepath
		name := filepath.Base(page)

		// Create am empty template set and use Funcs to register the above template.FuncMap for our custom template function
		// Parse the base template into template set
		ts, err := template.New(name).Funcs(functions).ParseFiles("client/ui/HTML/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Add any partials into the tempalte set
		ts, err = ts.ParseGlob("client/ui/HTML/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Adding the final 'page' template into the template set
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Adding parsed template set to the map, name of the template as the key
		cache[name] = ts

	}

	return cache, nil
}
