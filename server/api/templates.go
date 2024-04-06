package main

import "snippetbox.sting.net/internal/models"

// templateData type to act as a Holding structure for dynamic data to be passed onto templates.
type templateData struct {
	Snippet *models.Snippet
}
