CREATE TABLE admins (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
);