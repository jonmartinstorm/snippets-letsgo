package main

import "github.com/jonmartinstorm/snippets-letsgo/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
