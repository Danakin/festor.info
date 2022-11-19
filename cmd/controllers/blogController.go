package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Blog struct{}

func NewBlogController() *Blog {
	return &Blog{}
}

func (c *Blog) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/blog/index.page.tmpl"
	View(w, route, nil)
}

func (c *Blog) Show(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/blog/show.page.tmpl"
	fmt.Println()
	fmt.Println("PARAM TEST", chi.URLParam(r, "slug"), chi.URLParam(r, "articleID"))
	fmt.Println()
	View(w, route, nil)
}
