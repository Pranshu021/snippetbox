package main

import "snippetbox.sting.net/internal/models"

// templateData type to act as a Holding structure for dynamic data to be passed onto templates.
// Can add multiple structs and wrap it into one
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
