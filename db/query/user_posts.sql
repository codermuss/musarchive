-- name: InsertUserPost :one
INSERT INTO user_posts (user_id, blog_id) 
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUserPost :exec
DELETE FROM user_posts 
WHERE user_id = $1 AND blog_id = $2;

-- name: GetUserBlogs :many
SELECT b.id, b.title, b.summary, b.content, b.cover_image, b.created_at, b.updated_at, b.likes 
FROM blogs b
JOIN user_posts up ON up.blog_id = b.id
WHERE up.user_id = $1;

-- name: GetUserBlog :one
SELECT b.id, b.title, b.summary, b.content, b.cover_image, b.created_at, b.updated_at, b.likes 
FROM blogs b
JOIN user_posts up ON up.blog_id = b.id
WHERE up.user_id = $1 AND b.id = $2;