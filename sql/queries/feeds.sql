-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    )
RETURNING *;
-- name: GetFeeds :many
SELECT *
FROM feeds;
-- name: GetNextFeedToFetch :many
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT ?;
-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = time('now'),
    updated_at = time('now')
WHERE id = ?
RETURNING *;