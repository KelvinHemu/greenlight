package main

import "github.kelvinhemu.snippetbox/internal/models"

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}
