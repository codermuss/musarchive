// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: comments.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments 
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const getCommentsForPost = `-- name: GetCommentsForPost :many
SELECT id, post_id, user_id, content, created_at 
FROM comments 
WHERE post_id = $1 
ORDER BY created_at DESC
`

func (q *Queries) GetCommentsForPost(ctx context.Context, postID int32) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getCommentsForPost, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
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

const insertComment = `-- name: InsertComment :one
INSERT INTO comments (post_id, user_id, content, created_at) 
VALUES ($1, $2, $3, $4) 
RETURNING id, post_id, user_id, content, created_at
`

type InsertCommentParams struct {
	PostID    int32              `json:"post_id"`
	UserID    int32              `json:"user_id"`
	Content   string             `json:"content"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) InsertComment(ctx context.Context, arg InsertCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, insertComment,
		arg.PostID,
		arg.UserID,
		arg.Content,
		arg.CreatedAt,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}
