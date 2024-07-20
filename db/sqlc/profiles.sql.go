// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: profiles.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE user_id = $1
`

func (q *Queries) DeleteProfile(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteProfile, userID)
	return err
}

const getProfile = `-- name: GetProfile :one
SELECT user_id, bio, blog_count, like_count, follower_count 
FROM profiles 
WHERE user_id = $1
`

func (q *Queries) GetProfile(ctx context.Context, userID int32) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfile, userID)
	var i Profile
	err := row.Scan(
		&i.UserID,
		&i.Bio,
		&i.BlogCount,
		&i.LikeCount,
		&i.FollowerCount,
	)
	return i, err
}

const insertProfile = `-- name: InsertProfile :one
INSERT INTO profiles (user_id, bio, blog_count, like_count, follower_count) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING user_id, bio, blog_count, like_count, follower_count
`

type InsertProfileParams struct {
	UserID        int32       `json:"user_id"`
	Bio           pgtype.Text `json:"bio"`
	BlogCount     pgtype.Int4 `json:"blog_count"`
	LikeCount     pgtype.Int4 `json:"like_count"`
	FollowerCount pgtype.Int4 `json:"follower_count"`
}

func (q *Queries) InsertProfile(ctx context.Context, arg InsertProfileParams) (Profile, error) {
	row := q.db.QueryRow(ctx, insertProfile,
		arg.UserID,
		arg.Bio,
		arg.BlogCount,
		arg.LikeCount,
		arg.FollowerCount,
	)
	var i Profile
	err := row.Scan(
		&i.UserID,
		&i.Bio,
		&i.BlogCount,
		&i.LikeCount,
		&i.FollowerCount,
	)
	return i, err
}

const updateProfile = `-- name: UpdateProfile :one
UPDATE profiles 
    SET 
    bio = COALESCE($1,bio), 
    blog_count = COALESCE($2,blog_count), 
    like_count = COALESCE($3,like_count),
    follower_count = COALESCE($4,follower_count)
    WHERE user_id = $5
RETURNING user_id, bio, blog_count, like_count, follower_count
`

type UpdateProfileParams struct {
	Bio           pgtype.Text `json:"bio"`
	BlogCount     pgtype.Int4 `json:"blog_count"`
	LikeCount     pgtype.Int4 `json:"like_count"`
	FollowerCount pgtype.Int4 `json:"follower_count"`
	UserID        int32       `json:"user_id"`
}

func (q *Queries) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	row := q.db.QueryRow(ctx, updateProfile,
		arg.Bio,
		arg.BlogCount,
		arg.LikeCount,
		arg.FollowerCount,
		arg.UserID,
	)
	var i Profile
	err := row.Scan(
		&i.UserID,
		&i.Bio,
		&i.BlogCount,
		&i.LikeCount,
		&i.FollowerCount,
	)
	return i, err
}
