-- name: CreateApplication :one
INSERT INTO applications (
  name,
  source_text
) VALUES (
  $1, $2
)
RETURNING *;