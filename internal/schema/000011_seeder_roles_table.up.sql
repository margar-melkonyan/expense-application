SET TIMEZONE = 'Europe/Moscow';

INSERT INTO roles (title, display_title, permissions, created_at, updated_at)
VALUES ('user', 'User', '[
  "categories_read",
  "budgets_create",
  "budgets_read",
  "budgets_update",
  "budgets_delete",
  "users_read",
  "users_update"
]', now(), now());

INSERT INTO roles (title, display_title, permissions, created_at, updated_at)
VALUES ('admin', 'Admin', '[
  "categories_create",
  "categories_read",
  "categories_update",
  "categories_delete",
  "budgets_create",
  "budgets_read",
  "budgets_update",
  "budgets_delete",
  "users_read",
  "users_update",
  "roles_create",
  "roles_read",
  "roles_update",
  "roles_delete"
]', now(), now());