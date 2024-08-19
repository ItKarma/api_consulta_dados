package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Connectiondb() *sql.DB {
	db, err := sql.Open("sqlite", ".././db/cadsus.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	log.Println("Connected to the database successfully")
	return db
}
