package posts

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	PostStore *PostStore
}

func (h *Handler) AllPosts(c echo.Context) error {
	testFeeds := []int{1, 2, 3, 4}

	return c.JSON(http.StatusOK, testFeeds)
}

func (h *Handler) GetPost(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
