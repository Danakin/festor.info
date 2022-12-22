package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type Post struct {
	Id          string
	TypeId      int64
	Title       string
	Description string
	Image       string
	IsReleased  bool
	ReleasedAt  *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time

	Type *Type
	Tags []*Tag
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
	posts.id,
	posts.type_id,
	posts.title,
	posts.description,
	posts.released_at,
	posts.updated_at,
	types.title as type_title
FROM
	posts
JOIN
	types
	ON types.id = posts.type_id
WHERE
	posts.is_released = TRUE
	AND posts.released_at < NOW()
`
	if len(title) > 0 && typeId > 0 {
		query += `
	AND LOWER(posts.title) like $3 AND posts.type_id = $4
`
	} else if len(title) > 0 {
		query += `
	AND LOWER(posts.title) like $3		
`
	} else if typeId > 0 {
		query += `
	AND posts.type_id = $3
`
	}
	query += `
ORDER BY
	posts.updated_at DESC
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
		var post Post = Post{
			Type: &Type{},
		}
		err = rows.Scan(&post.Id, &post.TypeId, &post.Title, &post.Description, &post.ReleasedAt, &post.UpdatedAt, &post.Type.Title)
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

func (ps *PostService) Insert(post *Post) (*Post, error) {
	blogQuery := `SELECT id FROM types WHERE title = 'Blog'`
	row := ps.db.QueryRow(blogQuery)
	var typeId int
	err := row.Scan(&typeId)
	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO posts(type_id, title, description, is_released, released_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	row = ps.db.QueryRow(query, typeId, post.Title, post.Description, post.IsReleased, post.ReleasedAt)
	err = row.Scan(&post.Id)
	if err != nil {
		return nil, err
	}

	return post, nil
}
