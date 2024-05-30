package controls

import (
	"database/sql"
	"log"
	"user/db"
	"user/model"
)

func CreateUserTable(db *sql.DB) {
	// db := db.DbIn()
	query := `CREATE TABLE IF NOT EXISTS users(
		user_id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP
	)`
	res, err := db.Exec(query)
	if err != nil {
		log.Fatal("falied to create user table")
	}
	if res != nil {
		log.Fatal("table created")
	}
}

func InsertUser(model.User) {
	db := db.DbIn()
	defer db.Close()
	CreateUserTable(db)
}
