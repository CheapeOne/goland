package resolvers

import (
	"context"

	"github.com/cheapeone/goland/api/models"
)

func (r *mutationResolver) CreateFeed(ctx context.Context, input models.NewFeed) (*models.Feed, error) {
	panic("not implemented")
}

func (r *queryResolver) Feeds(ctx context.Context) ([]*models.Feed, error) {
	return r.fs.GetAll()
}
