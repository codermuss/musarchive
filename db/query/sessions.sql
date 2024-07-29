-- name: InsertSession :one
INSERT INTO sessions (id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetSession :one
SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at 
FROM sessions 
WHERE id = $1;

-- -- name: UpdateSession :one
-- UPDATE sessions 
-- SET user_id = $1, refresh_token = $2, user_agent = $3, client_ip = $4, is_blocked = $5, expires_at = $6, created_at = $7
-- WHERE id = $8
-- RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions 
WHERE id = $1;