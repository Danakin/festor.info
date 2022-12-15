package controllers

import (
	"net/http"
)

type Contact struct{}

func NewContactController() *Contact {
	return &Contact{}
}

func (c *Contact) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/contact.page.tmpl"
	view(w, r, route, nil)
}
