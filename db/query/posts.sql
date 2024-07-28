-- name: InsertPost :one
INSERT INTO posts (user_id, title, summary, content, cover_image, created_at, updated_at, likes) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING *;

-- name: GetPost :one
SELECT id, user_id, title, summary, content, cover_image, created_at, updated_at, likes 
FROM posts 
WHERE id = $1;

-- name: GetPosts :many
SELECT id, user_id, title, summary, content, cover_image, created_at, updated_at, likes 
FROM posts 
ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts 
    SET
    title = COALESCE(sqlc.narg(title),title), 
    summary = COALESCE(sqlc.narg(summary),summary), 
    content = COALESCE(sqlc.narg(content),content), 
    cover_image = COALESCE(sqlc.narg(cover_image),cover_image),  
    likes = COALESCE(sqlc.narg(likes),likes)
    WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts 
WHERE id = $1;