CREATE TABLE users (
   id serial PRIMARY KEY,
   email TEXT NOT NULL UNIQUE,
   hash TEXT NOT NULL
);
