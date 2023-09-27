-- +goose Up
CREATE TABLE gousers (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (hex(random()))
);
-- +goose Down
DROP TABLE gousers;