-- CREATE users(
--
-- );

CREATE TABLE budgets(
    id SERIAL PRIMARY KEY,
    type VARCHAR(255),
    amount BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(2048) UNIQUE,
    slug VARCHAR(2048),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE budget_category(
    budget_id INTEGER,
    category_id INTEGER
);