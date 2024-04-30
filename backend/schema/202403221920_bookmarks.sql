-- migrate:up

CREATE TABLE bookmarks
(
    id        SERIAL                      NOT NULL UNIQUE,
    user_id   INTEGER                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    area_id   INTEGER                     NOT NULL REFERENCES areas (id) ON DELETE CASCADE,
    room_id   INTEGER                     NOT NULL REFERENCES rooms (id) ON DELETE CASCADE,
    skill_id  INTEGER                     NOT NULL REFERENCES skills (id) ON DELETE CASCADE,
    ctime     TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW() at time zone 'utc'),
    completed BOOLEAN DEFAULT FALSE
);

-- migrate:down

DROP TABLE bookmarks;
