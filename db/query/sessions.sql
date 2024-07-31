-- name: InsertSession :one
INSERT INTO sessions (id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetSession :one
SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at 
FROM sessions 
WHERE id = $1;

-- name: UpdateSession :exec
UPDATE sessions 
SET is_blocked = $1
WHERE id = $2;

-- name: DeleteSession :exec
DELETE FROM sessions 
WHERE id = $1;