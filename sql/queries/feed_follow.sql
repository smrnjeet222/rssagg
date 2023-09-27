-- name: CreateFeedFollow :one
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (?, ?, ?, ?, ?)
RETURNING *;
-- name: GetFeedFollows :many 
SELECT *
FROM feeds_follows
WHERE user_id = ? ;
-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows
WHERE id = ?
    AND user_id = ?;