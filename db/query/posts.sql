-- name: InsertPost :one
INSERT INTO posts (user_id, title, content, cover_image, created_at, updated_at, likes) 
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetPost :one
SELECT id, user_id, title, content, cover_image, created_at, updated_at, likes 
FROM posts 
WHERE id = $1;

-- name: GetPosts :many
SELECT id, user_id, title, content, cover_image, created_at, updated_at, likes 
FROM posts 
ORDER BY id LIMIT $1 OFFSET $2;

-- name: GetPostsWithFilter :many
SELECT DISTINCT p.id, p.user_id, p.title, p.content, p.cover_image, p.created_at, p.updated_at, p.likes 
FROM posts p
LEFT JOIN post_categories pc ON p.id = pc.post_id
LEFT JOIN post_tags pt ON p.id = pt.post_id
WHERE 
    (array_length($3::integer[], 1) IS NULL OR pc.category_id = ANY($3)) AND
    (array_length($4::integer[], 1) IS NULL OR pt.tag_id = ANY($4))
ORDER BY p.id LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts 
    SET
    title = COALESCE(sqlc.narg(title),title), 
    content = COALESCE(sqlc.narg(content),content), 
    cover_image = COALESCE(sqlc.narg(cover_image),cover_image),  
    likes = COALESCE(sqlc.narg(likes),likes)
    WHERE id = sqlc.arg(id)
RETURNING *;

-- name: GetFollowedPosts :many
SELECT p.id, p.user_id, p.title, p.content, p.cover_image, p.created_at, p.updated_at, p.likes 
FROM posts p
JOIN user_followers f ON p.user_id = f.user_id
WHERE f.follower_id = $1
ORDER BY p.id LIMIT $2 OFFSET $3;

-- name: DeletePost :exec
DELETE FROM posts 
WHERE id = $1;