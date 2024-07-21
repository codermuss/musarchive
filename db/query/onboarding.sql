-- name: InsertOnboarding :one
INSERT INTO onboarding (image, title, description) 
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetOnboarding :one
SELECT id, image, title, description 
FROM onboarding
WHERE id = $1;

-- name: ListOnboarding :many
SELECT id, image, title, description 
FROM onboarding;

-- name: UpdateOnboarding :one
UPDATE onboarding
SET image = $1, title = $2, description = $3
WHERE id = $4
RETURNING *;

-- name: DeleteOnboarding :exec
DELETE FROM onboarding 
WHERE id = $1;