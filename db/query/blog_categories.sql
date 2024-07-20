-- name: InsertBlogCategory :one
INSERT INTO blog_categories (blog_id, category_id) 
VALUES ($1, $2)
RETURNING *;

-- name: GetCategoriesForBlog :many
SELECT c.id, c.name 
FROM categories c
JOIN blog_categories bc ON bc.category_id = c.id
WHERE bc.blog_id = $1;

-- name: DeleteBlogCategory :exec
DELETE FROM blog_categories 
WHERE blog_id = $1 AND category_id = $2;