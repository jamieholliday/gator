-- +goose Up
CREATE TABLE IF NOT EXISTS posts (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    title text NOT NULL,
    url text NOT NULL UNIQUE,
    description text NOT NULL,
    published_at timestamp NOT NULL,
    feed_id uuid NOT NULL REFERENCES feeds (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS posts;

