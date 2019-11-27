package resolvers

import (
	"context"

	"github.com/cheapeone/goland/api/models"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	post := &models.Post{Title: "green"}
	return []*models.Post{post}, nil
}
