package controllers

import (
	"net/http"
)

type Error struct{}

func NewErrorController() *Error {
	return &Error{}
}

func (c *Error) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/error.page.tmpl"
	view(w, r, route, nil)
}
