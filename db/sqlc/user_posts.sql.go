// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_posts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteUserPost = `-- name: DeleteUserPost :exec
DELETE FROM user_posts 
WHERE user_id = $1 AND post_id = $2
`

type DeleteUserPostParams struct {
	UserID int32 `json:"user_id"`
	PostID int32 `json:"post_id"`
}

func (q *Queries) DeleteUserPost(ctx context.Context, arg DeleteUserPostParams) error {
	_, err := q.db.Exec(ctx, deleteUserPost, arg.UserID, arg.PostID)
	return err
}

const getUserPost = `-- name: GetUserPost :one
SELECT b.id, b.title, b.summary, b.content, b.cover_image, b.created_at, b.updated_at, b.likes 
FROM posts b
JOIN user_posts up ON up.post_id = b.id
WHERE up.user_id = $1 AND b.id = $2
`

type GetUserPostParams struct {
	UserID int32 `json:"user_id"`
	ID     int32 `json:"id"`
}

type GetUserPostRow struct {
	ID         int32              `json:"id"`
	Title      string             `json:"title"`
	Summary    string             `json:"summary"`
	Content    string             `json:"content"`
	CoverImage pgtype.Text        `json:"cover_image"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	Likes      pgtype.Int4        `json:"likes"`
}

func (q *Queries) GetUserPost(ctx context.Context, arg GetUserPostParams) (GetUserPostRow, error) {
	row := q.db.QueryRow(ctx, getUserPost, arg.UserID, arg.ID)
	var i GetUserPostRow
	err := row.Scan(
		&i.ID,
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

const getUserPosts = `-- name: GetUserPosts :many
SELECT b.id, b.title, b.summary, b.content, b.cover_image, b.created_at, b.updated_at, b.likes 
FROM posts b
JOIN user_posts up ON up.post_id = b.id
WHERE up.user_id = $1
`

type GetUserPostsRow struct {
	ID         int32              `json:"id"`
	Title      string             `json:"title"`
	Summary    string             `json:"summary"`
	Content    string             `json:"content"`
	CoverImage pgtype.Text        `json:"cover_image"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	Likes      pgtype.Int4        `json:"likes"`
}

func (q *Queries) GetUserPosts(ctx context.Context, userID int32) ([]GetUserPostsRow, error) {
	rows, err := q.db.Query(ctx, getUserPosts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUserPostsRow{}
	for rows.Next() {
		var i GetUserPostsRow
		if err := rows.Scan(
			&i.ID,
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

const insertUserPost = `-- name: InsertUserPost :one
INSERT INTO user_posts (user_id, post_id) 
VALUES ($1, $2)
RETURNING user_id, post_id
`

type InsertUserPostParams struct {
	UserID int32 `json:"user_id"`
	PostID int32 `json:"post_id"`
}

func (q *Queries) InsertUserPost(ctx context.Context, arg InsertUserPostParams) (UserPost, error) {
	row := q.db.QueryRow(ctx, insertUserPost, arg.UserID, arg.PostID)
	var i UserPost
	err := row.Scan(&i.UserID, &i.PostID)
	return i, err
}
