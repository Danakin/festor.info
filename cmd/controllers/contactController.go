package controllers

import (
	"html/template"
	"net/http"
)

type Contact struct {
	TemplateCache map[string]*template.Template
}

func NewContactController(templateCache map[string]*template.Template) *Contact {
	return &Contact{
		TemplateCache: templateCache,
	}
}

func (c *Contact) Index(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/contact.page.tmpl"
	View(route, nil, w, c.TemplateCache)
}
