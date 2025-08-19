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
  SET source_text = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM applications
WHERE id = $1;