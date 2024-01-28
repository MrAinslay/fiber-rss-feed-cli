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

func commandCreate(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")

	switch splitString[0] {
	case "user":
		jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "password": "%s"}`, splitString[1], splitString[2]))
		bodyReader := bytes.NewReader(jsonBody)
		rsp, err := cfg.ApiClient.HttpClient.Post(fmt.Sprintf("%s/users", cfg.ApiClient.BaseURL), "application/json", bodyReader)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.User{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Printf("\n\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	case "feed":
		jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "url": "%s"}`, splitString[1], splitString[2]))
		bodyReader := bytes.NewReader(jsonBody)

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/feeds", cfg.ApiClient.BaseURL), bodyReader)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()
		defer req.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.Feed{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Printf("\n\nID: %s\nCreated At: %s\nUser ID: %s\nName: %s\nURL: %s\n\n", params.Id, params.CreatedAt, params.UserId, params.Name, params.URL)
	case "feed-follow":
		if len(splitString) < 2 {
			return errors.New("not enough arguments")
		}

		jsonBody := []byte(fmt.Sprintf(`{"feed_id": "%s"}`, splitString[1]))
		bodyReader := bytes.NewReader(jsonBody)
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/feed-follows", cfg.ApiClient.BaseURL), bodyReader)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()
		defer req.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.FeedFollow{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Printf("\nID: %s\nCreated At: %s\nUser ID: %s\nFeed ID: %s\n\n", params.Id, params.CreatedAt, params.UserId, params.FeedId)
	case "post-like":
		if len(splitString) < 2 {
			return errors.New("not enough arguments")
		}

		jsonBody := []byte(fmt.Sprintf(`{"post_id": "%s"}`, splitString[1]))
		bodyReader := bytes.NewReader(jsonBody)
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/post-likes", cfg.ApiClient.BaseURL), bodyReader)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()
		defer req.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.PostLike{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Printf("\nID: %s\nCreated At: %s\nUser ID: %s\nFeed ID: %s\n\n", params.Id, params.CreatedAt, params.UserId, params.PostId)
	}

	return nil
}
