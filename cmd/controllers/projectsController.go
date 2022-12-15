package controllers

import (
	"net/http"
)

type Project struct{}

func NewProjectController() *Project {
	return &Project{}
}

func (c *Project) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/projects.page.tmpl"
	view(w, r, route, nil)
}
