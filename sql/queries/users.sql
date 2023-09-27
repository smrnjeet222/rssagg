-- name: CreateUser :one
INSERT INTO gousers (id, created_at, updated_at, name, api_key)
VALUES (
        ?,
        ?,
        ?,
        ?,
        -- encode(sha256(random()::text::bytea), 'hex')
        hex(random())
    )
RETURNING *;
-- name: GetUserByAPIKey :one
SELECT *
FROM gousers
WHERE api_key = ?