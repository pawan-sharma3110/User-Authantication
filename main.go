package main

import (
	"user/db"
	"user/db/controls"
)

func main() {
	db := db.DbIn()
	controls.CreateUserTable(db)

	// r := router.UserRouts()
	// http.ListenAndServe(":8080", r)
}
