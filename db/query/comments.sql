-- name: InsertComment :one
INSERT INTO comments (post_id, user_id, content, created_at) 
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetCommentsForPost :many
SELECT id, post_id, user_id, content, created_at 
FROM comments 
WHERE post_id = $1 
ORDER BY created_at DESC;

-- name: DeleteComment :exec
DELETE FROM comments 
WHERE id = $1;