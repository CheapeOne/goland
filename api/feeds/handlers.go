package feeds

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	FeedStore *FeedStore
}

func (h *Handler) AllFeeds(c echo.Context) error {
	testFeeds := []int{1, 2, 3, 4}

	return c.JSON(http.StatusOK, testFeeds)
}

func (h *Handler) GetFeed(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
