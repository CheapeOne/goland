package feeds

import (
	"github.com/jmoiron/sqlx"
)

type FeedStore struct {
	db *sqlx.DB
}

func NewFeedStore(db *sqlx.DB) *FeedStore {
	return &FeedStore{db: db}
}

func (fs *FeedStore) FindAll() ([]Feed, int, error) {
	var feeds []Feed
	var count int
	return feeds, count, nil
}

func (fs *FeedStore) Find() {}
