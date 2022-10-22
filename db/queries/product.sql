-- name: CreateProduct :one
INSERT INTO products
    (sku, name, country)
VALUES
    ($1,$2, $3)
RETURNING *;