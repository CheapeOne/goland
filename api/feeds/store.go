package feeds

import "github.com/jinzhu/gorm"

type FeedStore struct {
	db *gorm.DB
}

func NewFeedStore(db *gorm.DB) *FeedStore {
	return &FeedStore{db: db}
}

func (fs *FeedStore) FindAll() ([]Feed, int, error) {
	var feeds []Feed
	var count int
	fs.db.Model(&feeds).Count(&count)
	fs.db.Find(&feeds)
	return feeds, count, nil
}

func (fs *FeedStore) Find()
