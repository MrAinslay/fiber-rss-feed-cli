package api

import (
	"time"

	"github.com/google/uuid"
)

type FeedFollow struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
	ErrorMsg  string    `json:"error"`
}
