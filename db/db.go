package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

func Migrate(db Database) error {
	query := `
    DROP TABLE IF EXISTS purchases;
    DROP TABLE IF EXISTS servises;
    DROP TABLE IF EXISTS users;
    DROP TABLE IF EXISTS orders;
        CREATE TABLE IF NOT EXISTS users(
        id SERIAL PRIMARY KEY,
        cash VARCHAR(100) NOT NULL ,
        reserve VARCHAR(100) NOT NULL DEFAULT 0
        );
        CREATE TABLE IF NOT EXISTS servises(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        cost VARCHAR(100) NOT NULL
        );
        CREATE TABLE IF NOT EXISTS orders(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) 
        );
        CREATE TABLE IF NOT EXISTS purchases(
            id SERIAL PRIMARY KEY,
            userid INTEGER REFERENCES users (id),
            serviseid INTEGER REFERENCES servises (id),
            orderid INTEGER REFERENCES orders (id),
            paid BOOLEAN DEFAULT false,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );
        INSERT INTO users(cash) VALUES (100);
        INSERT INTO users(cash) VALUES (99);
        
        INSERT INTO servises(name, cost) VALUES ('first', 100);
        INSERT INTO servises(name, cost) VALUES ('last', 99);
        
        INSERT INTO orders(name) VALUES ('123');
        INSERT INTO orders(name) VALUES ('23');
        INSERT INTO orders(name) VALUES ('3');            
        `
	_ = db.Conn.QueryRow(query)
	return nil
}
