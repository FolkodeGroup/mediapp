-- +goose Up
INSERT INTO users (username, email, password)
VALUES ('medico_prueba', 'medico@prueba.com', '$2b$10$QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm123456');

-- +goose Down
DELETE FROM users WHERE email = 'medico@prueba.com';
