package api

import (
	"github.com/cheapeone/goland/api/feeds"
	"github.com/cheapeone/goland/api/posts"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Api struct {
	FeedStore *feeds.FeedStore
	PostStore *posts.PostStore
}

func NewApi(db *gorm.DB) *Api {
	fs := feeds.NewFeedStore(db)
	ps := posts.NewPostStore(db)
	return &Api{FeedStore: fs, PostStore: ps}
}

func (api *Api) Register(router *echo.Group) {
	feedHandler := feeds.Handler{FeedStore: api.FeedStore}
	router.GET("/feeds", feedHandler.AllFeeds)
	router.GET("/feeds/{feedID}", feedHandler.GetFeed)

	postHandler := posts.Handler{PostStore: api.PostStore}
	router.GET("/posts", postHandler.AllPosts)
	router.GET("/posts.{postId}", postHandler.GetPost)
}
