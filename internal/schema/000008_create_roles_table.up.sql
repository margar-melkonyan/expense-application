CREATE TABLE IF NOT EXISTS roles
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR(100),
    display_title VARCHAR(255),
    permissions   JSONB,
    created_at    TIMESTAMP,
    updated_at    TIMESTAMP,
    deleted_at    TIMESTAMP DEFAULT NULL
);