package controllers

import (
	"html/template"
	"net/http"
)

type Homepage struct {
	TemplateCache map[string]*template.Template
}

func NewHomepageController(templateCache map[string]*template.Template) *Homepage {
	return &Homepage{
		TemplateCache: templateCache,
	}
}

func (c *Homepage) Index(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/index.page.tmpl"
	View(route, nil, w, c.TemplateCache)
}
