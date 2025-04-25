CREATE TABLE restaurants (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    address TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    description TEXT NOT NULL
);