-- name: GetAllCategories :many
SELECT
  id,
  name,
  description
FROM categories
WHERE id <> 'DEF';

-- name: GetCategoryById :one
SELECT
  id,
  name,
  description
FROM categories
WHERE id = $1;

-- name: AddCategory :exec
INSERT INTO categories (id, name, description)
VALUES ($1, $2, $3);

-- name: UpdateCategoryById :exec
UPDATE categories
SET
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description)
WHERE id = $1;

-- name: DeleteCategoryById :exec
DELETE FROM categories
WHERE id = $1;