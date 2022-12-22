package models

import (
	"database/sql"
	"time"
)

type Type struct {
	Id          int64
	Title       string
	Description string
	Icon        string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type TypeService struct {
	db *sql.DB
}

func (ts *TypeService) Get() ([]Type, error) {
	query := `
SELECT id, title
FROM TYPES;
	`
	rows, err := ts.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []Type
	for rows.Next() {
		var t Type = Type{}
		err = rows.Scan(&t.Id, &t.Title)
		if err != nil {
			return nil, err
		}

		types = append(types, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}
