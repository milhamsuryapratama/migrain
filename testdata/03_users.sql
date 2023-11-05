-- UP --
CREATE TABLE users (
    id BIGINT,
    name VARCHAR(255) NOT NULL DEFAULT '',
    age char(2) NULL DEFAULT NULL,
    address text NULL DEFAULT NULL,
    created_at DATETIME,
    updated_at DATETIME
);

-- DOWN --
DROP TABLE users;