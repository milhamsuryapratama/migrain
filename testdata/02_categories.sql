-- UP --
CREATE TABLE categories (
    id BIGINT,
    name VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

-- DOWN --
DROP TABLE categories;