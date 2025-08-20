-- name: CreateApplication :one
INSERT INTO applications (
  name,
  source_text
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetApplication :one
SELECT * FROM applications
WHERE id = $1 LIMIT 1;


-- name: ListApplications :many
SELECT * FROM applications
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: UpdateApplication :one
UPDATE applications
  SET translation_text = COALESCE(sqlc.narg(translation_text), translation_text)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteApplication :exec
DELETE FROM applications
WHERE id = $1;