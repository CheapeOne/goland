package resolvers

import (
	"context"

	"github.com/cheapeone/goland/api"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*api.Post, error) {
	post := &api.Post{Title: "green"}
	return []*api.Post{post}, nil

}
