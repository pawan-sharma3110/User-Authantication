package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbIn() *sql.DB {
	ConnStr := `host=localhost port=5432 user=postgres password=Pawan@2003 dbname=P-kart sslmode=disable`
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatalf("Error while exucting connection string %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error while conecting to database %v", err)
	}
	return db
}
