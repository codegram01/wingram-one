
BEGIN;

CREATE TABLE gram (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    parent_id INT,
    account_id INT,

    FOREIGN KEY (parent_id) REFERENCES "gram" (id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE
);

END;
