CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    password TEXT NOT NULL
);

CREATE TABLE token (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE
);

CREATE TABLE refresh_token (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL,
    token_id INT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE,
    FOREIGN KEY (token_id) REFERENCES token (id) ON DELETE CASCADE
);

CREATE TABLE profile (
    id SERIAL PRIMARY KEY,
    name TEXT,
    account_id INT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES "account" (id) ON DELETE CASCADE
);

CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    content TEXT,
    profile_id INT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (profile_id) REFERENCES "profile" (id) ON DELETE CASCADE
);