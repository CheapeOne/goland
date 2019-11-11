package feeds

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Feed struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	RssUrl      string
	SelfUrl     string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
