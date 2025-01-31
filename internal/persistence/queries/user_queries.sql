-- name: InsertUser :one
INSERT INTO users (id, name, email, password, created_at, created_by)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = $1;