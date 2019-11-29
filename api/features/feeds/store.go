package feeds

import (
	"github.com/cheapeone/goland/api/models"
	"github.com/jmoiron/sqlx"
)

type FeedStore struct {
	DB *sqlx.DB
}

func (fs *FeedStore) GetAll() ([]*models.Feed, error) {
	feed := &models.Feed{Title: "green"}
	return []*models.Feed{feed}, nil
}
