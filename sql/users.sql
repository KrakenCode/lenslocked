DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    first_name TEXT,
    last_name TEXT,
    age INT
);


INSERT INTO users (email, first_name, last_name, age) VALUES('example@example.com', 'ex-first', 'ex-last', 30)