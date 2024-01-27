-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS products (
    product_id CHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price FLOAT,
    stock_quantity INT
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back;
DROP TABLE products;
