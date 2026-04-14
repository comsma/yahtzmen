-- name: GetProduct :one
SELECT id, name, price, description, notes, dimensions, features, is_live, created_at
FROM products
WHERE id = ?;

-- name: GetAllProducts :many
SELECT id, name, price, description, notes, dimensions, features, is_live, created_at
FROM products;

-- name: CreateProduct :exec
INSERT INTO products (id, name, price, description, notes, dimensions, features, is_live)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateProduct :exec
UPDATE products
SET name = ?, price = ?, description = ?, notes = ?, dimensions = ?, features = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;

-- name: GetProductImages :many
SELECT id, product_id, object_key, sequence_number, created_at
FROM product_images
WHERE product_id = ?
ORDER BY sequence_number ASC;

-- name: CreateProductImage :exec
INSERT INTO product_images (id, product_id, object_key, sequence_number)
VALUES (?, ?, ?, ?);

-- name: GetAllOrders :many
SELECT id, ship_name, ship_add1, ship_add2, ship_add_postal, created_at
FROM orders;

-- name: GetOrder :one
SELECT id, ship_name, ship_add1, ship_add2, ship_add_postal, created_at
FROM orders
WHERE id = ?;

-- name: CreateOrder :exec
INSERT INTO orders (id, ship_name, ship_add1, ship_add2, ship_add_postal)
VALUES (?, ?, ?, ?, ?);
