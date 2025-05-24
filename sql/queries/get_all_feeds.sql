-- name: GetAllFeeds :many
SELECT
    feeds.*,
    users.name AS user_name
FROM
    feeds
    LEFT JOIN users ON feeds.user_id = users.id;

