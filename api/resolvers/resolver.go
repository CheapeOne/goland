package resolvers

import (
	"github.com/cheapeone/goland/api"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() api.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() api.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
