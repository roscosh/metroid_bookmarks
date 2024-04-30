-- migrate:up

CREATE TABLE areas
(
    id      SERIAL      NOT NULL UNIQUE,
    name_en VARCHAR(50) NOT NULL UNIQUE,
    name_ru VARCHAR(50) NOT NULL UNIQUE
);

-- migrate:down

DROP TABLE areas;