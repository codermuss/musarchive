-- name: InsertUserPost :one
INSERT INTO user_posts (user_id, post_id) 
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUserPost :exec
DELETE FROM user_posts 
WHERE user_id = $1 AND post_id = $2;

-- name: GetUserPosts :many
SELECT b.id, b.title, b.summary, b.content, b.cover_image, b.created_at, b.updated_at, b.likes 
FROM posts b
JOIN user_posts up ON up.post_id = b.id
WHERE up.user_id = $1;

-- name: GetUserPost :one
SELECT b.id, b.title, b.summary, b.content, b.cover_image, b.created_at, b.updated_at, b.likes 
FROM posts b
JOIN user_posts up ON up.post_id = b.id
WHERE up.user_id = $1 AND b.id = $2;