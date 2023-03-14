CREATE TABLE articles (
    id BIGINT,
    user_id BIGINT,
    title VARCHAR(255),
    content TEXT,
    created_at DATETIME,
    updated_at DATETIME
);