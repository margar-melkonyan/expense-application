CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    tg_id    BIGINT UNIQUE DEFAULT NULL,
    name     VARCHAR(255),
    email    VARCHAR(255),
    password VARCHAR(255)  DEFAULT NULL
);