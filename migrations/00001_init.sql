-- +goose Up
CREATE TABLE products
(
    id          VARCHAR(20)    NOT NULL PRIMARY KEY,
    name        VARCHAR(255)   NOT NULL,
    description TEXT,
    notes       TEXT,
    features    TEXT,
    dimensions  TEXT,
    created_at  TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price       DECIMAL(10, 2) NOT NULL,
    is_live     BOOLEAN        NOT NULL DEFAULT FALSE
);

CREATE TABLE product_images
(
    id              VARCHAR(20)  NOT NULL PRIMARY KEY,
    product_id      VARCHAR(20)  NOT NULL,
    object_key      VARCHAR(500) NOT NULL DEFAULT '',
    sequence_number INT          NOT NULL DEFAULT 0,
    created_at      TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE orders
(
    id              VARCHAR(20)  NOT NULL PRIMARY KEY,
    ship_name       VARCHAR(255) NOT NULL,
    ship_add1       VARCHAR(255) NOT NULL,
    ship_add2       VARCHAR(255) NOT NULL DEFAULT '',
    ship_add_postal VARCHAR(20)  NOT NULL,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE product_images;
DROP TABLE orders;
DROP TABLE products;
