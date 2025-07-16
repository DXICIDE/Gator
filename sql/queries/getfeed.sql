-- name: GetFeeds :many
SELECT feeds.name as feed_name, feeds.url, users.name as user_name FROM feeds
JOIN users ON feeds.user_id = users.id;