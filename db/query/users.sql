-- name: InsertUser :one
INSERT INTO users (username, password, full_name, email, avatar,role, birth_date, password_changed_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9) 
RETURNING *;

-- name: GetUser :one
SELECT id, username, password, full_name, email, avatar,role, birth_date, password_changed_at, created_at 
FROM users 
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users 
SET username = $1, password = $2, full_name = $3, email = $4, avatar = $5, role=$6, birth_date = $7, password_changed_at = $8, created_at = $9
WHERE id = $10
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;