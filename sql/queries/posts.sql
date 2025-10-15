-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
	gen_random_uuid(),
	now(),
	now(),
	$1,
	$2,
	$3,
	$4,
	$5
);

-- name: GetPostsForUser :many
SELECT *
FROM posts
WHERE feed_id IN (
	SELECT feed_follows.feed_id FROM feed_follows WHERE feed_follows.user_id = $1	
)
ORDER BY published_at DESC
LIMIT $2;


