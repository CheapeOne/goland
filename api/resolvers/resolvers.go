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
	fs *feeds.FeedService
	ps *posts.PostService
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() api.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() api.QueryResolver {
	return &queryResolver{r}
}

func GraphqlHandler(db *sqlx.DB) http.HandlerFunc {
	feedService := feeds.NewFeedService(db)
	postService := posts.NewPostService(db)
	resolver := &Resolver{
		fs: feedService,
		ps: postService,
	}

	return handler.GraphQL(api.NewExecutableSchema(api.Config{Resolvers: resolver}))
}
