CREATE TABLE IF NOT EXISTS invites
(
    id          serial PRIMARY KEY,
    user_id  INTEGER NOT NULL REFERENCES users(id),
    event_id  INTEGER NOT NULL REFERENCES events(id),
    type       VARCHAR(128),
    approval       VARCHAR(128),
    created_at  timestamptz NOT NULL DEFAULT Now(),
    modified_at timestamptz NOT NULL DEFAULT Now(),
    deleted_at timestamptz,
);
