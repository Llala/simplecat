-- name: CreateTranslationUnit :one
INSERT INTO translation_unit (
  application_id,
  source_unit_id,
  text
) VALUES (
  $1, $2, NULL
)
RETURNING *;

-- name: GetTranslationUnit :one
SELECT * FROM translation_unit
WHERE id = $1 LIMIT 1;


-- name: ListTranslationUnits :many
SELECT * FROM translation_unit
WHERE application_id = $1
ORDER BY id;


-- name: UpdateTranslationUnit :one

UPDATE translation_unit
SET
  source_unit_id = COALESCE(sqlc.narg(source_unit_id), source_unit_id),
  text = COALESCE(sqlc.narg(text), text)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTranslationUnit :exec
DELETE FROM translation_unit
WHERE id = $1;



