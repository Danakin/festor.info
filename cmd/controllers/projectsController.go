package controllers

import (
	"html/template"
	"net/http"
)

type Project struct {
	TemplateCache map[string]*template.Template
}

func NewProjectController(templateCache map[string]*template.Template) *Project {
	return &Project{
		TemplateCache: templateCache,
	}
}

func (c *Project) Index(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/projects.page.tmpl"
	View(route, nil, w, c.TemplateCache)
}
