package api

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
