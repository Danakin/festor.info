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
	view(w, r, route, nil)
}

func (c *Blog) Show(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/blog/show.page.tmpl"
	fmt.Println()
	fmt.Println("PARAM TEST", "slug", chi.URLParam(r, "slug"), "articleId", chi.URLParam(r, "articleID"))
	fmt.Println()

	data := &templateData{}
	data.Slug = chi.URLParam(r, "slug")
	// data.Slug = chi.URLParam(r, "slug")

	view(w, r, route, data)
}
