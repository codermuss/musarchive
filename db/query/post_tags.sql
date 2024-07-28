-- name: InsertPostTag :one
INSERT INTO post_tags (post_id, tag_id) 
VALUES ($1, $2)
RETURNING *;

-- name: GetTagsForPost :many
SELECT t.id, t.name 
FROM tags t
JOIN post_tags bt ON bt.tag_id = t.id
WHERE bt.post_id = $1;

-- name: DeletePostTag :exec
DELETE FROM post_tags 
WHERE post_id = $1 AND tag_id = $2;