DELETE
FROM user_roles
WHERE (role_id, user_id) IN
      (SELECT role_id, user_id
       FROM users
                INNER JOIN user_roles ON users.id = user_roles.user_id
                INNER JOIN roles ON user_roles.role_id = roles.id
       WHERE users.email = 'admin@admin.ru'
         AND roles.title = 'admin'
       LIMIT 1);

DELETE
FROM users
WHERE email = 'admin@admin.ru'