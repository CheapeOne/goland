package posts

import "github.com/jmoiron/sqlx"

type PostStore struct {
	db *sqlx.DB
}

func NewPostStore(db *sqlx.DB) *PostStore {
	return &PostStore{db: db}
}

func (ps *PostStore) GetAllPosts() {
}
