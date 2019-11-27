package feeds

import (
	"github.com/cheapeone/goland/api/models"
	"github.com/jmoiron/sqlx"
)

type FeedService struct {
	db *sqlx.DB
}

func NewFeedService(db *sqlx.DB) *FeedService {
	return &FeedService{db: db}
}

func (fs *FeedService) GetAll() ([]*models.Feed, error) {
	feed := &models.Feed{Title: "green"}
	return []*models.Feed{feed}, nil
}
