-- name: CreateFeedFollow :one
WITH feed_follows as(
	INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
	VALUES(
		gen_random_uuid(),
		now(),
		now(),
		$1,
		$2
	) RETURNING *) SELECT feed_follows.id as feed_id, feed_follows.created_at as feed_created_at, feed_follows.updated_at as feed_updated_at, users.name as user_name, feeds.name as feed_name 
		FROM feed_follows
		INNER JOIN users ON feed_follows.user_id = users.id
		INNER JOIN feeds ON feed_follows.feed_id = feeds.id;

-- name: SelectAllFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name
	FROM feed_follows
	INNER JOIN feeds ON feed_follows.feed_id = feeds.id
	WHERE feed_follows.user_id = $1;
