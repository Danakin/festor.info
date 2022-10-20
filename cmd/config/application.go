package config

import (
	"html/template"

	"github.com/danakin/festor.info/cmd/controllers"
)

type Application struct {
	TemplateCache map[string]*template.Template
	Controllers   *controllers.Controllers
}

func NewApplication(templatePath string) (*Application, error) {
	templateCache, err := newTemplateCache(templatePath)
	if err != nil {
		return nil, err
	}

	controllers := controllers.NewControllers(templateCache)

	return &Application{
		TemplateCache: templateCache,
		Controllers:   controllers,
	}, nil
}
