package config

import (
	"html/template"
	"io/fs"
	"path/filepath"
)

type templateData struct {
	// AuthenticatedUser *models.User
	CSRFToken   string
	CurrentYear int
	Flash       string
	// Form              *forms.Form
	// Snippet           *models.Snippet
	// Snippets          []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages := []string{}
	err := filepath.Walk(filepath.Join(dir, "pages"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		isTemplate, e := filepath.Match("*.page.tmpl", filepath.Base(path))
		if e != nil {
			return e
		}

		if isTemplate {
			pages = append(pages, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := page                         // strip the path, leave only filename *.page.tmpl
		ts, err := template.ParseFiles(page) // create template set that contains current page
		// ts, err := template.New(name).Funcs(functions).ParseFiles(page) // pass custom functions to the template with the current name
		if err != nil {
			return nil, err
		}

		// get all layout and partial files
		ts, err = ts.ParseGlob(filepath.Join(dir, "layouts", "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "partials", "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
