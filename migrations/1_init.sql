-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS products (
    product_id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS items (
    item_id CHAR(36) PRIMARY KEY,
    product_id VARCHAR(100) NOT NULL,
    price FLOAT,
    stock_quantity INT,
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE items;
DROP TABLE products;
