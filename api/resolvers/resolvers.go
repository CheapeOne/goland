package resolvers

import (
	"net/http"

	"github.com/cheapeone/goland/api"

	"github.com/99designs/gqlgen/handler"
	"github.com/cheapeone/goland/api/features/feeds"
	"github.com/cheapeone/goland/api/features/posts"
	"github.com/jmoiron/sqlx"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	fs *feeds.FeedStore
	ps *posts.PostStore
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() api.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() api.QueryResolver {
	return &queryResolver{r}
}

func NewGraphqlHandler(db *sqlx.DB) http.HandlerFunc {
	fs := &feeds.FeedStore{DB: db}
	ps := &posts.PostStore{DB: db}
	resolver := &Resolver{fs, ps}

	return handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: resolver}))
}
