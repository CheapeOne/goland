package posts

import "github.com/jmoiron/sqlx"

type PostService struct {
	db *sqlx.DB
}

func NewPostService(db *sqlx.DB) *PostService {
	return &PostService{db: db}
}
