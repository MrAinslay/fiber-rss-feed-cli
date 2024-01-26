package api

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HttpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		BaseURL: baseURL,
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}
