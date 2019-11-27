package resolvers

import (
	"context"

	"github.com/cheapeone/goland/api"
)

func (r *mutationResolver) CreateFeed(ctx context.Context, input api.NewFeed) (*api.Feed, error) {
	panic("not implemented")
}

func (r *queryResolver) Feeds(ctx context.Context) ([]*api.Feed, error) {
	panic("not implemented")
}
