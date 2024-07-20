-- name: InsertBlog :one
INSERT INTO blogs (user_id, title, summary, content, cover_image, created_at, updated_at, likes) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
RETURNING *;

-- name: GetBlog :one
SELECT id, user_id, title, summary, content, cover_image, created_at, updated_at, likes 
FROM blogs 
WHERE id = $1;

-- name: UpdateBlog :one
UPDATE blogs 
    SET
    title = COALESCE(sqlc.narg(title),title), 
    summary = COALESCE(sqlc.narg(summary),summary), 
    content = COALESCE(sqlc.narg(content),content), 
    cover_image = COALESCE(sqlc.narg(cover_image),cover_image),  
    likes = COALESCE(sqlc.narg(likes),likes)
    WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs 
WHERE id = $1;