package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Printf("\n\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	case "feeds":
		rsp, err := cfg.ApiClient.HttpClient.Get(fmt.Sprintf("%s/feeds", cfg.ApiClient.BaseURL))
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)

		params := []api.Feed{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params[0].ErrorMsg != "" {
			return errors.New(params[0].ErrorMsg)
		}

		for _, feed := range params {
			fmt.Printf("\n\nID: %s\nCreated At: %s\nUser ID: %s\nName: %s\nURL: %s\n\n", feed.Id, feed.CreatedAt, feed.UserId, feed.Name, feed.URL)
		}
	case "posts":
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/posts", cfg.ApiClient.BaseURL), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := []api.Post{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params[0].ErrorMsg != "" {
			return errors.New(params[0].ErrorMsg)
		}

		for _, post := range params {
			fmt.Printf("\n\nID: %s\nCreated At: %s\nTitle: %s\nURL: %s\nDescription: %s\nPublished At: %s\nFeed ID: %s\n", post.Id, post.CreatedAt, post.Title, post.URL, post.Description, post.PublishedAt, post.FeedID)
		}
	case "feed-follows":
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/feed-follows", cfg.ApiClient.BaseURL), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := []api.Feed{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params[0].ErrorMsg != "" {
			return errors.New(params[0].ErrorMsg)
		}

		for _, feed := range params {
			fmt.Printf("\n\nID: %s\nCreated At: %s\nURL: %s\nName: %s\n\n", feed.Id, feed.CreatedAt, feed.Name, feed.URL)
		}
	case "post-likes":
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/post-likes", cfg.ApiClient.BaseURL), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := []api.Post{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params[0].ErrorMsg != "" {
			return errors.New(params[0].ErrorMsg)
		}

		for _, post := range params {
			fmt.Printf("\n\nID: %s\nCreated At: %s\nTitle: %s\nURL: %s\nDescription: %s\nPublished At: %s\nFeed ID: %s\n", post.Id, post.CreatedAt, post.Title, post.URL, post.Description, post.PublishedAt, post.FeedID)
		}
	case "feed-id":
		if len(splitString) < 2 {
			return errors.New("not enough arguments")
		}

		rsp, err := cfg.ApiClient.HttpClient.Get(fmt.Sprintf("%s/feeds/%s", cfg.ApiClient.BaseURL, splitString[1]))
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

		fmt.Printf("\n\nID: %s\nCreated At: %s\nURL: %s\nName: %s\n\n", params.Id, params.CreatedAt, params.Name, params.URL)
	}
	return nil
}
