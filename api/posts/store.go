package posts

import "github.com/jinzhu/gorm"

type PostStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{db: db}
}

func (ps *PostStore) GetAllPosts() {
}
