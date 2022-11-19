package controllers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/danakin/festor.info/ui"
)

type Controllers struct {
	Homepage   *Homepage
	Blog       *Blog
	Technology *Technology
	Contact    *Contact
	CV         *CV
	Project    *Project
}

func NewControllers() *Controllers {
	return &Controllers{
		Homepage:   NewHomepageController(),
		Blog:       NewBlogController(),
		Technology: NewTechnologyController(),
		Contact:    NewContactController(),
		CV:         NewCVController(),
		Project:    NewProjectController(),
	}
}

func View(w http.ResponseWriter, view string, data any) {
	ts, err := template.
		New(view).
		ParseFS(
			ui.EmbeddedFiles,
			"templates/layouts/app.layout.tmpl",
			"templates/partials/*.partial.tmpl",
			view,
		)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buf, "app", data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	buf.WriteTo(w)
}
