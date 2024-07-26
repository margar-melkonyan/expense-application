CREATE TABLE IF NOT EXISTS user_roles
(
    user_id bigint,
    role_id bigint,

    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);