-- name: InsertProfile :one
INSERT INTO profiles (user_id, bio, blog_count, like_count, follower_count) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetProfile :one
SELECT user_id, bio, blog_count, like_count, follower_count 
FROM profiles 
WHERE user_id = $1;

-- name: UpdateProfile :one
UPDATE profiles 
    SET 
    bio = COALESCE(sqlc.narg(bio),bio), 
    blog_count = COALESCE(sqlc.narg(blog_count),blog_count), 
    like_count = COALESCE(sqlc.narg(like_count),like_count),
    follower_count = COALESCE(sqlc.narg(follower_count),follower_count)
    WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE user_id = $1;