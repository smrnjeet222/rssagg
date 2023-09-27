-- name: CreatePost :one
INSERT INTO posts (
        id,
        created_at,
        updated_at,
        title,
        description,
        published_at,
        url,
        feed_id
    )
VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    )
RETURNING *;
-- name: GetPostsForUser :many
SELECT posts.*
from posts
    JOIN feeds_follows ON posts.feed_id = feeds_follows.feed_id
WHERE feeds_follows.user_id = ?
ORDER BY posts.published_at DESC
LIMIT ?;