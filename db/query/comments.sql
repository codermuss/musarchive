-- name: InsertComment :one
INSERT INTO comments (blog_id, user_id, content, created_at) 
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetCommentsForBlog :many
SELECT id, blog_id, user_id, content, created_at 
FROM comments 
WHERE blog_id = $1 
ORDER BY created_at DESC;

-- name: DeleteComment :exec
DELETE FROM comments 
WHERE id = $1;