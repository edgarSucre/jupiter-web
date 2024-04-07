-- name: CreateUser :one
INSERT INTO users (
    admin, name, email, password
) VALUES (
    ?, ?, ?, ?
)
RETURNING *;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;