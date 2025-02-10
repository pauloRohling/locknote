-- name: InsertNote :one
INSERT INTO notes (id, title, content, created_at, created_by, updated_at, updated_by)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: FindNoteByID :one
SELECT *
FROM notes
WHERE id = $1
  AND created_by = $2;

-- name: FindNotesByUser :many
SELECT *
FROM notes
WHERE created_by = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateNote :one
UPDATE notes
SET title = $3, content = $4, updated_at = $5, updated_by = $6
WHERE id = $1
  AND created_by = $2
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1
  AND created_by = $2;

-- name: DeleteNotesByUser :exec
DELETE FROM notes
WHERE created_by = $1;