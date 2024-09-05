
BEGIN;

CREATE TABLE refresh_token (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    token_id INT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES token (id) ON DELETE CASCADE
);

END;
