-- UP --
CREATE TABLE roles (
    id BIGINT,
    name VARCHAR(255) NOT NULL DEFAULT '',
    created_at DATETIME,
    updated_at DATETIME
);

-- DOWN --
DROP TABLE roles;