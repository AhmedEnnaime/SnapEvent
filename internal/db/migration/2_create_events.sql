CREATE TABLE IF NOT EXISTS events
(
    id          serial PRIMARY KEY,
    eventDate   DATE NOT NULL,
    time        VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    city        VARCHAR(128) NOT NULL,
    location    VARCHAR(128) NOT NULL,
    status      VARCHAR(128) NOT NULL,
    creator_id  INTEGER NOT NULL REFERENCES users(id),
    created_at  timestamptz NOT NULL DEFAULT Now(),
    updated_at timestamptz NOT NULL DEFAULT Now(),
    deleted_at timestamptz 
);
