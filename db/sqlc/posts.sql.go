// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts 
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePost, id)
	return err
}

const getFollowedPosts = `-- name: GetFollowedPosts :many
SELECT p.id, p.user_id, p.title, p.content, p.cover_image, p.created_at, p.updated_at, p.likes 
FROM posts p
JOIN user_followers f ON p.user_id = f.user_id
WHERE f.follower_id = $1
ORDER BY p.id LIMIT $2 OFFSET $3
`

type GetFollowedPostsParams struct {
	FollowerID int32 `json:"follower_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetFollowedPosts(ctx context.Context, arg GetFollowedPostsParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, getFollowedPosts, arg.FollowerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
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

const getPost = `-- name: GetPost :one
SELECT id, user_id, title, content, cover_image, created_at, updated_at, likes 
FROM posts 
WHERE id = $1
`

func (q *Queries) GetPost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CoverImage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Likes,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, user_id, title, content, cover_image, created_at, updated_at, likes 
FROM posts 
ORDER BY id LIMIT $1 OFFSET $2
`

type GetPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, getPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
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

const getPostsWithFilter = `-- name: GetPostsWithFilter :many
SELECT DISTINCT p.id, p.user_id, p.title, p.content, p.cover_image, p.created_at, p.updated_at, p.likes 
FROM posts p
LEFT JOIN post_categories pc ON p.id = pc.post_id
LEFT JOIN post_tags pt ON p.id = pt.post_id
WHERE 
    (array_length($3::integer[], 1) IS NULL OR pc.category_id = ANY($3)) AND
    (array_length($4::integer[], 1) IS NULL OR pt.tag_id = ANY($4))
ORDER BY p.id LIMIT $1 OFFSET $2
`

type GetPostsWithFilterParams struct {
	Limit   int32   `json:"limit"`
	Offset  int32   `json:"offset"`
	Column3 []int32 `json:"column_3"`
	Column4 []int32 `json:"column_4"`
}

func (q *Queries) GetPostsWithFilter(ctx context.Context, arg GetPostsWithFilterParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, getPostsWithFilter,
		arg.Limit,
		arg.Offset,
		arg.Column3,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
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

const insertPost = `-- name: InsertPost :one
INSERT INTO posts (user_id, title, content, cover_image, created_at, updated_at, likes) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING id, user_id, title, content, cover_image, created_at, updated_at, likes
`

type InsertPostParams struct {
	UserID     pgtype.Int4 `json:"user_id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	CoverImage pgtype.Text `json:"cover_image"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Likes      int32       `json:"likes"`
}

func (q *Queries) InsertPost(ctx context.Context, arg InsertPostParams) (Post, error) {
	row := q.db.QueryRow(ctx, insertPost,
		arg.UserID,
		arg.Title,
		arg.Content,
		arg.CoverImage,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Likes,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CoverImage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Likes,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts 
    SET
    title = COALESCE($1,title), 
    content = COALESCE($2,content), 
    cover_image = COALESCE($3,cover_image),  
    likes = COALESCE($4,likes)
    WHERE id = $5
RETURNING id, user_id, title, content, cover_image, created_at, updated_at, likes
`

type UpdatePostParams struct {
	Title      pgtype.Text `json:"title"`
	Content    pgtype.Text `json:"content"`
	CoverImage pgtype.Text `json:"cover_image"`
	Likes      pgtype.Int4 `json:"likes"`
	ID         int32       `json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost,
		arg.Title,
		arg.Content,
		arg.CoverImage,
		arg.Likes,
		arg.ID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Content,
		&i.CoverImage,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Likes,
	)
	return i, err
}
