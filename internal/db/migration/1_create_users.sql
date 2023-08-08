CREATE TABLE IF NOT EXISTS users
(
    id          serial PRIMARY KEY,
    name        VARCHAR(128) NOT NULL,
    birthday    DATE NOT NULL, 
    email       VARCHAR(128) NOT NULL,
    password    VARCHAR(128) NOT NULL,
    gender      VARCHAR(128) NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT Now(),
    modified_at timestamptz NOT NULL DEFAULT Now()
);
