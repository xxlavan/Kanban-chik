package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseURL string) *sql.DB {
	const maxAttempts = 5
	const delay = 2 * time.Second

	var db *sql.DB
	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		log.Printf("connecting to database (attempt %d/%d)", attempt, maxAttempts)

		db, err = sql.Open("pgx", databaseURL)
		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			log.Println("successfully connected to database")
			return db
		}

		log.Printf("database not ready: %v", err)
		time.Sleep(delay)
	}

	log.Fatalf("failed to connect to database after %d attempts", maxAttempts)
	return nil
}
