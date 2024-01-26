package main

import (
	"time"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
)

func main() {
	client := api.NewClient(5 * time.Second)
	cfg := ApiConfig{
		ApiClient: client,
		ApiKey:    "",
	}
	startRepl(&cfg)
}
