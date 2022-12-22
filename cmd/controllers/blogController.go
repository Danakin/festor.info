package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

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

// TODO: Pagination Build Helper
func (c *Blog) Index(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/blog/index.page.tmpl"
	tplData := templateData{}

	title := strings.TrimSpace(r.URL.Query().Get("title"))
	typeId, err := strconv.Atoi(r.URL.Query().Get("type_id"))
	if err != nil {
		typeId = 0
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 25
	}

	tplData.Search = struct {
		Title  string
		TypeId int
	}{
		Title:  title,
		TypeId: typeId,
	}

	types, err := c.TypeService.Get()
	if err != nil {
		// TODO: Abortcontroller
		fmt.Println("%w", err)
		return
	}

	offset := (page - 1) * limit
	posts, total, err := c.PostService.Paginate(limit, offset, title, typeId)
	if err != nil {
		fmt.Println("%w", err)
		return
	}
	tplData.Pagination = &pagination{
		Page:    page,
		PerPage: limit,
		Total:   total,
	}

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

func (c *Blog) Create(w http.ResponseWriter, r *http.Request) {
	route := "templates/pages/blog/create.page.tmpl"
	tplData := &templateData{}
	tplData.Data = struct {
		Title       string
		Description string
		IsReleased  bool
		ReleasedAt  string
	}{
		Title:       "",
		Description: "",
		IsReleased:  false,
		ReleasedAt:  "",
	}

	view(w, r, route, tplData)
}

func (c *Blog) Store(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	isReleased := r.Form.Has("is_released")
	releasedAt := r.FormValue("released_at")

	// TODO: validation
	// TODO: CSRF

	// 2022-12-22T13:05
	t, err := time.Parse("2006-01-02T15:04", releasedAt)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	post := &models.Post{
		Title:       title,
		Description: description,
		IsReleased:  isReleased,
		ReleasedAt:  &t,
	}
	post, err = c.PostService.Insert(post)
	if err != nil {
		fmt.Println("%w", err)
		return
	}

	fmt.Println("Post after update...", post)

	http.Redirect(w, r, "/blog", http.StatusFound)
}
