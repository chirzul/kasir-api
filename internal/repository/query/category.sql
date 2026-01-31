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

-- name: UpdateCategoryById :execresult
UPDATE categories
SET
  name = $1,
  description = $2,
  updated_at = now()
WHERE id = $3;

-- name: DeleteCategoryById :execresult
DELETE FROM categories
WHERE id = $1;