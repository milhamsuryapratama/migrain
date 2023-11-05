-- UP --
CREATE TABLE products (
    id BIGINT,
    name VARCHAR(255) NOT NULL DEFAULT '',
    price double precision NOT NULL DEFAULT 0,
    stock int NOT NULL DEFAULT 0,
    created_at DATETIME,
    updated_at DATETIME
);

-- DOWN --
DROP TABLE products;