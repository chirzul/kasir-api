-- name: GetAllProducts :many
SELECT
  p.id,
  p.name,
  p.price,
  p.stock,
  c.name AS category_name,
  c.description AS category_description
FROM products p
JOIN categories c ON p.category_id = c.id;

-- name: GetProductById :one
SELECT
  p.id,
  p.name,
  p.price,
  p.stock,
  c.name AS category_name,
  c.description AS category_description
FROM products p
JOIN categories c ON p.category_id = c.id
WHERE p.id = $1;

-- name: AddProduct :exec
INSERT INTO products (name, price, stock, category_id)
VALUES ($1, $2, $3, $4);

-- name: UpdateProductById :execresult
UPDATE products
SET
  name = $1,
  price = $2,
  stock = $3,
  category_id = $4,
  updated_at = now()
WHERE id = $5;

-- name: DeleteProductById :execresult
DELETE FROM products
WHERE id = $1;