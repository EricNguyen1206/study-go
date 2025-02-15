package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDatabase() (*sql.DB, error) {
	if db == nil {
		var err error

		// TODO: change to env variable
		dsn := "root:password@tcp(127.0.0.1:3306)/test"
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			return nil, err
		}

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Fatalf("Database ping failed: %v", err)
			return nil, err
		}
		log.Println("Database connection established")
	}
	return db, nil
}

func ConnectDB() *sql.DB {
	return db
}
