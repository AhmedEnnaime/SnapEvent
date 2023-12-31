CREATE TABLE IF NOT EXISTS user_event
(
    id          serial PRIMARY KEY,
    user_id  INTEGER NOT NULL REFERENCES users(id),
    event_id  INTEGER NOT NULL REFERENCES events(id),
    created_at  timestamptz NOT NULL DEFAULT Now(),
    updated_at timestamptz NOT NULL DEFAULT Now(),
    deleted_at timestamptz,
);
