CREATE TABLE IF NOT EXISTS accounts (
    id bigserial PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT UNIQUE
)
