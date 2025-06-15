-- name: CreateUser :one
INSERT INTO users (id, name, role, created_at, updated_at, pw_hash)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE name = $1;

-- name: DropRows :exec
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

-- name: ListUsers :many
SELECT name FROM users;
