-- name: InsertPostCategory :one
INSERT INTO post_categories (post_id, category_id) 
VALUES ($1, $2)
RETURNING *;

-- name: GetCategoriesForPost :many
SELECT c.id, c.name 
FROM categories c
JOIN post_categories bc ON bc.category_id = c.id
WHERE bc.post_id = $1;

-- name: DeletePostCategory :exec
DELETE FROM post_categories 
WHERE post_id = $1 AND category_id = $2;