package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/danakin/festor.info/cmd/models"
	"github.com/go-chi/chi/v5"
)

type Blog struct {
	TypeService *models.TypeService
	PostService *models.PostService
}

func NewBlogController(services *models.Services) *Blog {
	return &Blog{
		TypeService: services.TypeService,
		PostService: services.PostService,
	}
}

func (c *Blog) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/blog/index.page.tmpl"
	tplData := templateData{}

	title := strings.TrimSpace(r.URL.Query().Get("title"))
	tagId, err := strconv.Atoi(r.URL.Query().Get("tag_id"))
	if err != nil {
		tagId = 0
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	tplData.Search = struct {
		Title string
		TagID int
	}{
		Title: title,
		TagID: tagId,
	}
	tplData.Pagination = &pagination{
		Page: &page,
	}

	types, err := c.TypeService.Get()
	if err != nil {
		// TODO: Abortcontroller
		fmt.Println("%w", err)
		return
	}

	posts, err := c.PostService.Get()

	tplData.Data = struct {
		Types []models.Type
		Posts []models.Post
	}{
		Types: types,
		Posts: posts,
	}
	fmt.Printf("%+v", tplData)

	view(w, r, route, &tplData)
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
