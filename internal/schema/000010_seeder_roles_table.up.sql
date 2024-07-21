SET TIMEZONE='Europe/Moscow';

INSERT INTO roles (title, display_title, permissions, created_at, updated_at)
VALUES ('user', 'User', '{
  "permissions": [
    "category_read",
    "budget_create",
    "budget_read",
    "budget_update",
    "budget_delete",
    "user_read",
    "user_update"
  ]
}', now(), now());

INSERT INTO roles (title, display_title, permissions, created_at, updated_at)
VALUES ('admin', 'Admin', '{
  "permissions": [
    "category_create",
    "category_read",
    "category_update",
    "category_delete",
    "budget_create",
    "budget_read",
    "budget_update",
    "budget_delete",
    "user_read",
    "user_update",
    "role_create",
    "role_read",
    "role_update",
    "role_delete"
  ]
}', now(), now());