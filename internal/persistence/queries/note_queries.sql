-- name: InsertNote :one
INSERT INTO notes (id, title, content, created_at, created_by)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: FindNoteByID :one
SELECT *
FROM notes
WHERE id = $1
  AND created_by = $2;

-- name: FindNotesByUser :many
SELECT *
FROM notes
WHERE created_by = $1
ORDER BY $2
LIMIT $3
OFFSET $4;