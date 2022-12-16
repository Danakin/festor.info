package config

import (
	"database/sql"

	"github.com/danakin/festor.info/cmd/models"
)

type Application struct {
	DB *sql.DB

	Services *models.Services
}

func NewApplication(db *sql.DB) (*Application, error) {
	services := models.NewServices(db)

	return &Application{
		DB:       db,
		Services: services,
	}, nil
}
