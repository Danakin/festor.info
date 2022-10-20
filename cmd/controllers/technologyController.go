package controllers

import (
	"html/template"
	"net/http"
)

type Technology struct {
	TemplateCache map[string]*template.Template
}

func NewTechnologyController(templateCache map[string]*template.Template) *Technology {
	return &Technology{
		TemplateCache: templateCache,
	}
}

func (c *Technology) Index(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/technologies.page.tmpl"
	View(route, nil, w, c.TemplateCache)
}
