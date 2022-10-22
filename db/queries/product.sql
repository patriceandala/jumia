-- name: CreateProduct :one
INSERT INTO Products
    (sku, name, country)
VALUES
    ($1,$2, $3)
RETURNING *;