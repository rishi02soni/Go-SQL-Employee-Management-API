package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		position TEXT,
		salary INTEGER
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
