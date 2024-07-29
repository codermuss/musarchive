-- name: InsertTag :one
INSERT INTO tags (name) 
VALUES ($1) 
RETURNING *;

-- name: GetTag :one
SELECT id, name 
FROM tags 
WHERE id = $1;

-- name: GetTags :many
SELECT id, name 
FROM tags;

-- name: UpdateTag :one
UPDATE tags 
SET name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tags 
WHERE id = $1;