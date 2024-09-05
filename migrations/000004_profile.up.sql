
BEGIN;

CREATE TABLE profile (
    id SERIAL PRIMARY KEY,
    name TEXT,
    account_id INT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE
);

END;
