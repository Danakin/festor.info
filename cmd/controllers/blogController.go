package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Blog struct {
	TemplateCache map[string]*template.Template
}

func NewBlogController(templateCache map[string]*template.Template) *Blog {
	return &Blog{
		TemplateCache: templateCache,
	}
}

func (c *Blog) Index(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/blog/index.page.tmpl"
	View(route, nil, w, c.TemplateCache)
}

func (c *Blog) Show(w http.ResponseWriter, r *http.Request) {
	route := "ui/templates/pages/blog/show.page.tmpl"
	fmt.Println()
	fmt.Println("PARAM TEST", chi.URLParam(r, "slug"), chi.URLParam(r, "articleID"))
	fmt.Println()
	View(route, nil, w, c.TemplateCache)
}
