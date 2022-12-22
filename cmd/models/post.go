package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Post struct {
	Id          string
	TypeId      uint64
	Title       string
	Description string
	Image       string
	IsReleased  bool
	ReleasedAt  *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type PostService struct {
	db *sql.DB
}

// TODO: query type, pagination, limit, order by column?
func (ps *PostService) Get() ([]Post, error) {
	query := `
SELECT
	id,
	type_id,
	title,
	description,
	updated_at
FROM
	posts
WHERE
	is_released = TRUE
	AND released_at < NOW()
ORDER BY
	updated_at DESC
`

	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.TypeId, &post.Title, &post.Description, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (ps *PostService) Paginate(limit int, offset int, title string, typeId int) ([]Post, int, error) {
	// TODO: Refactor to Build Search Query Because this is Ugly AF
	var total int
	countQuery := `
	SELECT count(*)
	FROM posts
	WHERE
		is_released = TRUE
		AND released_at < NOW()
	`
	var countRow *sql.Row
	if len(title) > 0 && typeId > 0 {
		countQuery += "AND LOWER(title) like $1 AND type_id = $2"
		countRow = ps.db.QueryRow(countQuery, fmt.Sprintf("%%%s%%", strings.ToLower(title)), typeId)
	} else if len(title) > 0 && typeId == 0 {
		countQuery += "AND LOWER(title) like $1"
		countRow = ps.db.QueryRow(countQuery, fmt.Sprintf("%%%s%%", strings.ToLower(title)))
	} else if len(title) == 0 && typeId > 0 {
		countQuery += "AND type_id = $1"
		countRow = ps.db.QueryRow(countQuery, typeId)
	} else {
		countRow = ps.db.QueryRow(countQuery)
	}
	if err := countRow.Scan(&total); err != nil {
		return nil, -1, err
	}

	query := `
SELECT
	id,
	type_id,
	title,
	description,
	updated_at
FROM
	posts
WHERE
	is_released = TRUE
	AND released_at < NOW()
`
	if len(title) > 0 && typeId > 0 {
		query += `
	AND LOWER(title) like $3 AND type_id = $4
`
	} else if len(title) > 0 {
		query += `
	AND LOWER(title) like $3		
`
	} else if typeId > 0 {
		query += `
	AND type_id = $3
`
	}
	query += `
ORDER BY
	updated_at DESC
LIMIT $1
OFFSET $2
`

	var rows *sql.Rows
	var err error
	if len(title) > 0 && typeId > 0 {
		rows, err = ps.db.Query(query, limit, offset, fmt.Sprintf("%%%s%%", strings.ToLower(title)), typeId)
	} else if len(title) > 0 {
		rows, err = ps.db.Query(query, limit, offset, fmt.Sprintf("%%%s%%", strings.ToLower(title)))
	} else if typeId > 0 {
		rows, err = ps.db.Query(query, limit, offset, typeId)
	} else {
		rows, err = ps.db.Query(query, limit, offset)
	}
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.TypeId, &post.Title, &post.Description, &post.UpdatedAt)
		if err != nil {
			return nil, -1, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, -1, err
	}

	return posts, total, nil
}
