package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.kelvinhemu.snippetbox/internal/models"
)

type templateData struct {
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	CurrentYear     int
	Form            any
	Flash           string
	IsAuthenticated bool
}

// humanDate formats a given time.Time into a human-readable string "DD MMM YYYY at HH:MM".
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// functions provide a mapping of custom template functions, including a "humanDate" function for formatting dates.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache
func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		// Parse the base template file into a template set.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

// render
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	files := []string{
		filepath.Join(uiDir, "html", "base.html"),
		filepath.Join(uiDir, "html", "partials", "nav.html"),
		filepath.Join(uiDir, "html", "pages", page),
	}

	ts, err := template.New("").Funcs(functions).ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
