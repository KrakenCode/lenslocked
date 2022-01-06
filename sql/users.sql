DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    first_name TEXT,
    last_name TEXT,
    age INT
);


INSERT INTO users (email, first_name, last_name, age) VALUES ('user1@lenslocked.com', 'user1', '', 99);

INSERT INTO users (email, first_name, last_name, age) VALUES ('user2@lenslocked.com', 'user2', '', 99);