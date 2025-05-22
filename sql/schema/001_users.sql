-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;

