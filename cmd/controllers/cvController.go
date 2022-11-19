package controllers

import (
	"net/http"
)

type CV struct{}

func NewCVController() *CV {
	return &CV{}
}

func (c *CV) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/cv.page.tmpl"
	View(w, route, nil)
}
