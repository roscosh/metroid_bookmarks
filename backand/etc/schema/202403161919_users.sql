-- migrate:up

CREATE TABLE users
(
    id       SERIAL       NOT NULL UNIQUE,
    name     VARCHAR(255) NOT NULL,
    login    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN      NOT NULL DEFAULT FALSE
);

-- migrate:down

DROP TABLE users;