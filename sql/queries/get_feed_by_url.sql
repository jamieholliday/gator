-- name: GetFeedByUrl :one
SELECT
    feeds.*
FROM
    feeds
WHERE
    feeds.url = $1;

