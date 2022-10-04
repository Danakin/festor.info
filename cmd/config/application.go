package config

import "html/template"

type Application struct {
	TemplateCache map[string]*template.Template
}

func NewApplication(templatePath string) (*Application, error) {
	templateCache, err := newTemplateCache(templatePath)
	if err != nil {
		return nil, err
	}

	return &Application{
		TemplateCache: templateCache,
	}, nil
}
