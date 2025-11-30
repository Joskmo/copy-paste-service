-- name: CreateNote :exec
INSERT INTO notes (id, content, created_at, expires_at)
VALUES ($1, $2, $3, $4);

-- name: GetNoteByID :one
SELECT id, content, created_at, expires_at
FROM notes
WHERE id = $1 AND expires_at > NOW();

-- name: DeleteNote :exec
DELETE FROM notes WHERE id = $1;

-- name: DeleteExpiredNotes :execrows
DELETE FROM notes WHERE expires_at <= NOW();

-- name: CountNotes :one
SELECT COUNT(*) FROM notes;

