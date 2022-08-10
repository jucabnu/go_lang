package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connection := "host=localhost user=postgres password=root dbname=juliano_loja sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func PingDb() {
	db := ConnectDatabase()
	defer db.Close()
}
