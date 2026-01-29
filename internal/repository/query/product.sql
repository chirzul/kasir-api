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

-- name: UpdateProductById :exec
UPDATE products
SET
  name = COALESCE(sqlc.narg(name), name),
  price = COALESCE(sqlc.narg(price), price),
  stock = COALESCE(sqlc.narg(stock), stock)
WHERE id = $1;

-- name: DeleteProductById :exec
DELETE FROM products
WHERE id = $1;