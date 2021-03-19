-- name: FindUserByID :one
SELECT * FROM users
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: FindUserByUUID :one
SELECT * FROM users
WHERE uuid = $1 AND deleted_at IS NULL LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND deleted_at IS NULL LIMIT 1;