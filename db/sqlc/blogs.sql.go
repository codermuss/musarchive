// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: blogs.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs 
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT id, user_id, title, summary, content, cover_image, created_at, updated_at, likes 
FROM blogs 
WHERE id = $1
`

func (q *Queries) GetBlog(ctx context.Context, id int32) (Blog, error) {
	row := q.db.QueryRow(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Content,
		&i.CoverImage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Likes,
	)
	return i, err
}

const getBlogs = `-- name: GetBlogs :many
SELECT id, user_id, title, summary, content, cover_image, created_at, updated_at, likes 
FROM blogs 
ORDER BY id LIMIT $1 OFFSET $2
`

type GetBlogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetBlogs(ctx context.Context, arg GetBlogsParams) ([]Blog, error) {
	rows, err := q.db.Query(ctx, getBlogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blog{}
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Summary,
			&i.Content,
			&i.CoverImage,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Likes,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertBlog = `-- name: InsertBlog :one
INSERT INTO blogs (user_id, title, summary, content, cover_image, created_at, updated_at, likes) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING id, user_id, title, summary, content, cover_image, created_at, updated_at, likes
`

type InsertBlogParams struct {
	UserID     pgtype.Int4        `json:"user_id"`
	Title      string             `json:"title"`
	Summary    string             `json:"summary"`
	Content    string             `json:"content"`
	CoverImage pgtype.Text        `json:"cover_image"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	Likes      pgtype.Int4        `json:"likes"`
}

func (q *Queries) InsertBlog(ctx context.Context, arg InsertBlogParams) (Blog, error) {
	row := q.db.QueryRow(ctx, insertBlog,
		arg.UserID,
		arg.Title,
		arg.Summary,
		arg.Content,
		arg.CoverImage,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Likes,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Content,
		&i.CoverImage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Likes,
	)
	return i, err
}

const updateBlog = `-- name: UpdateBlog :one
UPDATE blogs 
    SET
    title = COALESCE($1,title), 
    summary = COALESCE($2,summary), 
    content = COALESCE($3,content), 
    cover_image = COALESCE($4,cover_image),  
    likes = COALESCE($5,likes)
    WHERE id = $6
RETURNING id, user_id, title, summary, content, cover_image, created_at, updated_at, likes
`

type UpdateBlogParams struct {
	Title      pgtype.Text `json:"title"`
	Summary    pgtype.Text `json:"summary"`
	Content    pgtype.Text `json:"content"`
	CoverImage pgtype.Text `json:"cover_image"`
	Likes      pgtype.Int4 `json:"likes"`
	ID         int32       `json:"id"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error) {
	row := q.db.QueryRow(ctx, updateBlog,
		arg.Title,
		arg.Summary,
		arg.Content,
		arg.CoverImage,
		arg.Likes,
		arg.ID,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Summary,
		&i.Content,
		&i.CoverImage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Likes,
	)
	return i, err
}
