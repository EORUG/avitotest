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
    DROP TABLE IF EXISTS USERS;
    DROP TABLE IF EXISTS SEGMENTS;
    DROP TABLE IF EXISTS USER_SEGMENTS;
    DROP TABLE IF EXISTS LOG;
		
		CREATE TABLE IF NOT EXISTS USERS (
		userID SERIAL PRIMARY KEY
		);

		CREATE TABLE IF NOT EXISTS SEGMENTS (
		segmentID SERIAL PRIMARY KEY,
		segmentName varchar(100) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS USER_SEGMENTS (
		userID INTEGER REFERENCES USERS (userID),
		segmentID integer REFERENCES SEGMENTS (segmentID) ON DELETE CASCADE,
		TTL timestamp 
		);

		CREATE TABLE IF NOT EXISTS LOG (
		userID integer,
		segmentName varchar(100) NOT NULL,
		operation varchar(100) NOT NULL,
		datatime timestamp DEFAULT CURRENT_TIMESTAMP
		);

		CREATE OR REPLACE FUNCTION LOGERNEW()
		  RETURNS TRIGGER 
		  LANGUAGE PLPGSQL
		  AS
		$$
		DECLARE
		    _segmentName varchar(100);
		    _operation varchar(100) := TG_ARGV[0]::varchar(100);
		BEGIN
		  	SELECT s.segmentName
		  	INTO _segmentName
		  	FROM SEGMENTS s
		  	WHERE s.segmentID = NEW.segmentID;
			
			INSERT INTO LOG(userID,segmentName,operation,datatime)
		    VALUES(NEW.userID,_segmentName,_operation,now());
		  
		  RETURN NEW;
		END;		    
		$$
		    
		CREATE OR REPLACE FUNCTION LOGEROLD()
		  RETURNS TRIGGER 
		  LANGUAGE PLPGSQL
		  AS
		$$
		DECLARE
		    _segmentName varchar(100);
		    _operation varchar(100) := TG_ARGV[0]::varchar(100);
		BEGIN
		  	SELECT s.segmentName
		  	INTO _segmentName
		  	FROM SEGMENTS s
		  	WHERE s.segmentID = OLD.segmentID;
			
			INSERT INTO LOG(userID,segmentName,operation,datatime)
		    VALUES(OLD.userID,_segmentName,_operation,now());
		  
		  RETURN OLD;
		END;
		$$
		    
		CREATE TRIGGER segment_insert
		  BEFORE INSERT 
		  ON USER_SEGMENTS
		  FOR EACH ROW
		  EXECUTE PROCEDURE LOGERNEW('Добавилось'::varchar(100));

		CREATE TRIGGER segment_update
		  BEFORE delete 
		  ON USER_SEGMENTS
		  FOR EACH ROW
		  EXECUTE PROCEDURE LOGEROLD('Удалилось'::varchar(100)); 
        `
	_ = db.Conn.QueryRow(query)
	return nil
}
