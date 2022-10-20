package controllers

import (
	"html/template"
	"net/http"
)

type CV struct {
	TemplateCache map[string]*template.Template
}

func NewCVController(templateCache map[string]*template.Template) *CV {
	return &CV{
		TemplateCache: templateCache,
	}
}

func (c *CV) Index(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/cv.page.tmpl"
	View(route, nil, w, c.TemplateCache)
}
