package controllers

import (
	"bytes"
	"html/template"
	"net/http"
)

type Controllers struct {
	Homepage   *Homepage
	Blog       *Blog
	Technology *Technology
	Contact    *Contact
	CV         *CV
	Project    *Project
}

func NewControllers(templateCache map[string]*template.Template) *Controllers {
	return &Controllers{
		Homepage:   NewHomepageController(templateCache),
		Blog:       NewBlogController(templateCache),
		Technology: NewTechnologyController(templateCache),
		Contact:    NewContactController(templateCache),
		CV:         NewCVController(templateCache),
		Project:    NewProjectController(templateCache),
	}
}

func View(template string, data *interface{}, w http.ResponseWriter, templateCache map[string]*template.Template) {
	ts, ok := templateCache[template]
	if !ok {
		panic("not found")
		return
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, data)
	if err != nil {
		panic(err)
		return
	}
	buf.WriteTo(w)
}
