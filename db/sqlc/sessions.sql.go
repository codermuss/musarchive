// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sessions.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const deleteSession = `-- name: DeleteSession :exec

DELETE FROM sessions 
WHERE id = $1
`

// -- name: UpdateSession :one
// UPDATE sessions
// SET user_id = $1, refresh_token = $2, user_agent = $3, client_ip = $4, is_blocked = $5, expires_at = $6, created_at = $7
// WHERE id = $8
// RETURNING *;
func (q *Queries) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSession, id)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at 
FROM sessions 
WHERE id = $1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRow(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const insertSession = `-- name: InsertSession :one
INSERT INTO sessions (id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type InsertSessionParams struct {
	ID           uuid.UUID `json:"id"`
	UserID       int32     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (q *Queries) InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, insertSession,
		arg.ID,
		arg.UserID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
