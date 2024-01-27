package api

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uuid.UUID `json:"user_id"`
	PostId    uuid.UUID `json:"post_id"`
	ErrorMsg  string    `json:"error"`
}
