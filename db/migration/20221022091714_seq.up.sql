CREATE TABLE IF NOT EXISTS products
(
    "sku"     varchar PRIMARY KEY,
    "name"    varchar NOT NULL,
    "country" varchar NOT NULL,
    "stock"   integer NOT NULL DEFAULT 0
);
INSERT INTO products (sku, name, country, stock)
VALUES
    ('9befa247cd11','Chung PLC Table', 'eg',2),
    ('da8ef851e075','Turner-Payne Soft Bike', 'ci', 12),
    ('37d1fdc2cecb','Garcia, Jones and Murphy For repair Cotton Mouse', 'eg', 4),
    ('bb9eaf46d842','James and Sons Car', 'ma', 23);