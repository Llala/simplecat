-- name: CreateSourceUnit :one
INSERT INTO source_unit (
  application_id,
  translation_unit_id,
  text
) VALUES (
  $1, NULL, $2
)
RETURNING *;

-- name: GetSourceUnit :one
SELECT * FROM source_unit
WHERE id = $1 LIMIT 1;


-- name: ListSourceUnits :many
SELECT * FROM source_unit
WHERE application_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: UpdateSourceUnit :one

UPDATE source_unit
SET
  translation_unit_id = COALESCE(sqlc.narg(translation_unit_id), translation_unit_id),
  text = COALESCE(sqlc.narg(text), text)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteSourceUnit :exec
DELETE FROM source_unit
WHERE id = $1;

-- name: ListSourceUnitJoin :many
SELECT
    source_unit.id as source_unit_id,
    source_unit.text as source_text,
    translation_unit.text as translation_text,
    translation_unit.id as translation_unit_id
FROM
    source_unit 
INNER JOIN
    translation_unit ON source_unit.id = translation_unit.source_unit_id
WHERE source_unit.application_id = $1
ORDER BY source_unit.id
LIMIT $2
OFFSET $3;

-- name: ListSourceUnitJoinNoLimit :many
SELECT
    source_unit.id as source_unit_id,
    source_unit.text as source_text,
    translation_unit.text as translation_text,
    translation_unit.id as translation_unit_id
FROM
    source_unit 
INNER JOIN
    translation_unit ON source_unit.id = translation_unit.source_unit_id
WHERE source_unit.application_id = $1
ORDER BY source_unit.id;






