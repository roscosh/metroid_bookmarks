-- migrate:up

CREATE TABLE photos
(
    id          SERIAL       NOT NULL UNIQUE,
    bookmark_id INTEGER      NOT NULL REFERENCES bookmarks (id) ON DELETE CASCADE,
    path        varchar(250) NOT NULL UNIQUE
);

-- migrate:down

DROP TABLE photos;
