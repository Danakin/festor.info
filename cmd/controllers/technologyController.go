package controllers

import (
	"net/http"
)

type Technology struct{}

func NewTechnologyController() *Technology {
	return &Technology{}
}

func (c *Technology) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/technologies.page.tmpl"
	view(w, r, route, nil)
}
