package config

import (
	"html/template"
)

type Application struct {
	TemplateCache map[string]*template.Template
}

func NewApplication() (*Application, error) {
	// if err != nil {
	// 	return nil, err
	// }

	return &Application{}, nil
}
