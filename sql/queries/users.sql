-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
	gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: SelectUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: DeleteAll :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;
