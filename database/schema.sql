CREATE TABLE IF NOT EXISTS Users
(
    id                      SERIAL PRIMARY KEY,
    name                    TEXT NOT NULL,
    email                   TEXT NOT NULL,
    password                TEXT NOT NULL,
    created_at              TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    favorite_congressperson JSON NOT NULL DEFAULT '{}'
);