-- name: InsertUserFollower :one
INSERT INTO user_followers (user_id, follower_id) 
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUserFollower :exec
DELETE FROM user_followers 
WHERE user_id = $1 AND follower_id = $2;

-- name: GetFollowersOfUser :many
SELECT u.id, u.username, u.full_name, u.email, u.avatar, u.birth_date, u.password_changed_at, u.created_at
FROM users u
JOIN user_followers uf ON uf.follower_id = u.id
WHERE uf.user_id = $1;

-- name: GetFollowingUsers :many
SELECT u.id, u.username, u.full_name, u.email, u.avatar, u.birth_date, u.password_changed_at, u.created_at
FROM users u
JOIN user_followers uf ON uf.user_id = u.id
WHERE uf.follower_id = $1;
