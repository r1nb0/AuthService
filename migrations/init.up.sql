CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    nickname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_email ON users (email);