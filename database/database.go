package database

import (
	"github.com/jmoiron/sqlx"
)

func Connect() *sqlx.DB {
	db := &sqlx.DB{}
	return db
}
