// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, hashed_password, full_name, email, avatar, role, birth_date, is_email_verified, password_changed_at, created_at 
FROM users 
WHERE username = $1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Avatar,
		&i.Role,
		&i.BirthDate,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users (username, hashed_password, full_name, email, avatar,role, birth_date, password_changed_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9) 
RETURNING id, username, hashed_password, full_name, email, avatar, role, birth_date, is_email_verified, password_changed_at, created_at
`

type InsertUserParams struct {
	Username          string      `json:"username"`
	HashedPassword    string      `json:"hashed_password"`
	FullName          string      `json:"full_name"`
	Email             string      `json:"email"`
	Avatar            pgtype.Text `json:"avatar"`
	Role              string      `json:"role"`
	BirthDate         pgtype.Date `json:"birth_date"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRow(ctx, insertUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.Avatar,
		arg.Role,
		arg.BirthDate,
		arg.PasswordChangedAt,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Avatar,
		&i.Role,
		&i.BirthDate,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET username = $1, hashed_password = $2, full_name = $3, email = $4, avatar = $5, role=$6, birth_date = $7, is_email_verified=$8,password_changed_at = $9, created_at = $10
WHERE id = $11
RETURNING id, username, hashed_password, full_name, email, avatar, role, birth_date, is_email_verified, password_changed_at, created_at
`

type UpdateUserParams struct {
	Username          string      `json:"username"`
	HashedPassword    string      `json:"hashed_password"`
	FullName          string      `json:"full_name"`
	Email             string      `json:"email"`
	Avatar            pgtype.Text `json:"avatar"`
	Role              string      `json:"role"`
	BirthDate         pgtype.Date `json:"birth_date"`
	IsEmailVerified   bool        `json:"is_email_verified"`
	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
	ID                int32       `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.Avatar,
		arg.Role,
		arg.BirthDate,
		arg.IsEmailVerified,
		arg.PasswordChangedAt,
		arg.CreatedAt,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Avatar,
		&i.Role,
		&i.BirthDate,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
