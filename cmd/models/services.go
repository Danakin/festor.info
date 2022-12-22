package models

import "database/sql"

type Services struct {
	TypeService *TypeService
	PostService *PostService
	TagService  *TagService
}

func NewServices(db *sql.DB) *Services {
	return &Services{
		TypeService: &TypeService{db: db},
		PostService: &PostService{db: db},
		TagService:  &TagService{db: db},
	}
}
