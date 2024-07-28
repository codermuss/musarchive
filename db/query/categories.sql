-- name: InsertCategory :one
INSERT INTO categories (name) 
VALUES ($1) 
RETURNING *;


-- name: GetCategory :one
SELECT id, name 
FROM categories 
WHERE id = $1;


-- name: GetCategories :many
SELECT id, name 
FROM categories;

-- name: UpdateCategory :one
UPDATE categories 
SET name = $1
WHERE id = $2
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories 
WHERE id = $1;