package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "dev",
		Database: "lenslocked",
	}

	// connection string can also use this format: "postgres://postgres:dev@localhost:5432/lenslocked"
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panicln(err)
	}

	_, err = db.Exec(`
		DROP TABLE IF EXISTS customers; 

		CREATE TABLE customers(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		DROP TABLE IF EXISTS orders;

		CREATE TABLE orders(
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);

		ALTER TABLE orders ADD FOREIGN KEY (user_id) REFERENCES customers(id);
	`)

	_, err = db.Exec(`INSERT INTO customers(name, email) VALUES ($1, $2)`, "test-name", "test@example.com")
}
