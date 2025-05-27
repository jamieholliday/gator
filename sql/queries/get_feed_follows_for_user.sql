-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feeds_name
FROM
    feed_follows
    LEFT JOIN users ON feed_follows.user_id = users.id
    LEFT JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE
    feed_follows.user_id = $1;

