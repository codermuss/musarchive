-- name: InsertUser :one
INSERT INTO users (username, hashed_password, full_name, email, avatar,role, birth_date, password_changed_at, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8,$9) 
RETURNING *;

-- name: GetUser :one
SELECT * 
FROM users 
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users 
SET 
    username = COALESCE(sqlc.narg(username), username), 
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password), 
    full_name = COALESCE(sqlc.narg(full_name), full_name), 
    email = COALESCE(sqlc.narg(email), email), 
    avatar = COALESCE(sqlc.narg(avatar), avatar), 
    role = COALESCE(sqlc.narg(role), role), 
    birth_date = COALESCE(sqlc.narg(birth_date), birth_date), 
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified), 
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at), 
    created_at = COALESCE(sqlc.narg(created_at), created_at)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;