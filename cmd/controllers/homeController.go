package controllers

import (
	"net/http"
)

type Homepage struct{}

func NewHomepageController() *Homepage {
	return &Homepage{}
}

func (c *Homepage) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/index.page.tmpl"
	view(w, r, route, nil)
}
