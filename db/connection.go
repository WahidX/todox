package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nimda"
	dbname   = "todox"
)

func ConnectPostgres() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	DB = *db
}
