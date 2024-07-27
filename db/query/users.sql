-- name: InsertUser :one
INSERT INTO users (username, password, full_name, email, avatar, birth_date, password_changed_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING *;

-- name: GetUser :one
SELECT id, username, password, full_name, email, avatar, birth_date, password_changed_at, created_at 
FROM users 
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users 
SET username = $1, password = $2, full_name = $3, email = $4, avatar = $5, birth_date = $6, password_changed_at = $7, created_at = $8
WHERE id = $9
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;