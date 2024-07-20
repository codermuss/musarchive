-- name: InsertBlogTag :one
INSERT INTO blog_tags (blog_id, tag_id) 
VALUES ($1, $2)
RETURNING *;

-- name: GetTagsForBlog :many
SELECT t.id, t.name 
FROM tags t
JOIN blog_tags bt ON bt.tag_id = t.id
WHERE bt.blog_id = $1;

-- name: DeleteBlogTag :exec
DELETE FROM blog_tags 
WHERE blog_id = $1 AND tag_id = $2;