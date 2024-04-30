-- migrate:up

CREATE TABLE users
(
    id       SERIAL       NOT NULL UNIQUE,
    name     VARCHAR(255) NOT NULL,
    login    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE INDEX idx_users_login_password ON users (login, password);
CREATE INDEX idx_users_name_login ON users (name, login);

-- migrate:down

DROP TABLE users;