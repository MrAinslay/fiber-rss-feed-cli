package api

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
