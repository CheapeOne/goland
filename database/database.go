package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
