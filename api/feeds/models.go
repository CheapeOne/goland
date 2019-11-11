package feeds

import (
	"time"
)

type Feed struct {
	ID          int
	Title       string
	Description string
	RssUrl      string
	SelfUrl     string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
