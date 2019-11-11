package feeds

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	FeedStore *FeedStore
}

func (h *Handler) AllFeeds(c echo.Context) error {
	feeds, _, err := h.FeedStore.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, feeds)
}

func (h *Handler) GetFeed(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
