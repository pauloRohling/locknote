-- name: InsertNote :one
INSERT INTO notes (id, title, content, created_at, created_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;