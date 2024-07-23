SET TIMEZONE = 'Europe/Moscow';

INSERT INTO users (name, email, password, created_at, updated_at)
VALUES ('Admin',
        'admin@admin.ru',
        '$2a$10$xPf1gBp8Ppp1EgGOvJncx.9rx3uDLd.uZ0rwHdNv4i2uBLLdWKqTm', -- mQ-Nn2$01£PJ
        now(),
        now());