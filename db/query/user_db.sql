-- name: CreateUser :one
INSERT INTO users (
    admin, name, username, password
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: GetUserByUserName :one
SELECT *
FROM users
WHERE username = ?;