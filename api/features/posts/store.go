package posts

import "github.com/jmoiron/sqlx"

type PostStore struct {
	DB *sqlx.DB
}
