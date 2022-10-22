-- name: CreateProduct :one
INSERT INTO products
    (sku, name, country)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProduct :one
SELECT *
FROM products
WHERE sku = $1
LIMIT 1;

-- name: ConsumeProduct :one
UPDATE products
SET stock = stock - $1
WHERE sku = $2
  AND stock > $1
RETURNING *;