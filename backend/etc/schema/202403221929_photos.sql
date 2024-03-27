-- migrate:up

CREATE TABLE photos
(
    id          SERIAL      NOT NULL UNIQUE,
    Name        VARCHAR(25) NOT NULL,
    bookmark_id INTEGER     NOT NULL REFERENCES bookmarks (id) ON DELETE CASCADE
);

-- migrate:down

DROP TABLE photos;
