-- name: GetFeedFollowsForUser :many
SELECT feeds.name, users.name AS user_name FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id 
INNER JOIN users ON feed_follows.user_id = users.id
WHERE users.id = $1;