
BEGIN;

CREATE TABLE  account (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    password TEXT NOT NULL
);

END;
