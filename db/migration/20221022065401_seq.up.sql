CREATE TABLE IF NOT EXISTS products
(
    "sku"     varchar PRIMARY KEY,
    "name"    varchar NOT NULL,
    "country" varchar NOT NULL
);
-- INSERT INTO "Products" (sku, name, country)
-- VALUES
--     ('9befa247cd11','Chung PLC Table', 'eg'),
--     ('da8ef851e075','Turner-Payne Soft Bike', 'ci'),
--     ('37d1fdc2cecb','Garcia, Jones and Murphy For repair Cotton Mouse', 'eg'),
--     ('bb9eaf46d842','James and Sons Car', 'ma');