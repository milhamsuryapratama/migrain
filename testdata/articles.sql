-- UP --
CREATE TABLE articles (
    id BIGINT,
    user_id BIGINT,
    title VARCHAR(255),
    content TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE articles_test (
    id BIGINT,
    user_id BIGINT,
    title VARCHAR(255),
    content TEXT,
    created_at DATETIME,
    updated_at DATETIME
);

-- DOWN --
DROP TABLE articles;
DROP TABLE articles_test;