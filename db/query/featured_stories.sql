-- name: InsertFeaturedStory :one
INSERT INTO featured_stories (post_id, featured_date) 
VALUES ($1, $2) 
RETURNING *;

-- name: GetFeaturedStory :one
SELECT id, post_id, featured_date 
FROM featured_stories 
WHERE id = $1;

-- name: UpdateFeaturedStory :one
UPDATE featured_stories 
SET post_id = $1, featured_date = $2
WHERE id = $3
RETURNING *;

-- name: DeleteFeaturedStory :exec
DELETE FROM featured_stories 
WHERE id = $1;