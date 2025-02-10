-- name: InsertUser :one
INSERT INTO users (id, name, email, password, created_at, created_by, updated_at, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: FindUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: FindUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, updated_at = $3, updated_by = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;