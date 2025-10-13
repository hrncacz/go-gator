-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
	gen_random_uuid(),
	now(),
	now(),
	$1,
	$2,
	$3
)
RETURNING *;

-- name: SelectAllFeedsWithUsername :many
SELECT 
	feeds.name,
	feeds.url,
	users.name as username
FROM feeds
	INNER JOIN users
	ON feeds.user_id = users.id;

-- name: SelectFeedByUrl :one
SELECT *
	FROM feeds
	WHERE feeds.url = $1;
