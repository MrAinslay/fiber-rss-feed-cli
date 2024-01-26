package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
)

func commandGet(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")
	switch splitString[0] {
	case "user":
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/users", cfg.ApiClient.BaseURL), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		resp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		params := api.User{}
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErroMsg != "" {
			return errors.New(params.ErroMsg)
		}
		log.Printf("\n\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	case "feed":
		rsp, err := cfg.ApiClient.HttpClient.Get(fmt.Sprintf("%s/feeds", cfg.ApiKey))
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)

		params := api.Feed{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		log.Printf("\n\nID: %s\nCreated At: %s\nURL: %s\nName: %s\n\n", params.Id, params.CreatedAt, params.Name, params.URL)
	case "post":
		rsp, err := cfg.ApiClient.HttpClient.Get(fmt.Sprintf("%s/posts", cfg.ApiClient.BaseURL))
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := api.Post{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		log.Printf("\n\nID: %s\nCreated At: %s\nTitle: %s\nURL: %s\nDescription: %s\nPublished At: %s\nFeed ID: %s\n", params.Id, params.CreatedAt, params.Title, params.URL, params.Description, params.PublishedAt, params.FeedID)
	}
	return nil
}
