package api

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"updated_at"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt string    `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
	ErrorMsg    string    `json:"error"`
}
