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

func commandDelete(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")

	switch splitString[0] {
	case "user":
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users", cfg.ApiClient.BaseURL), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Auhtorization", cfg.ApiKey)

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		log.Println(params.Message)
	case "feed":
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/feeds/%s", cfg.ApiClient.BaseURL, splitString[1]), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", cfg.ApiKey)

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		log.Println(params.Message)
	case "feed_follow":
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/feed-follows/%s", cfg.ApiClient.BaseURL, splitString[1]), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", cfg.ApiKey)

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		log.Printf(params.Message)
	case "post-like":
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/post-likes/%s", cfg.ApiClient.BaseURL, splitString[1]), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", cfg.ApiKey)

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		log.Println(params.Message)
	}
	return nil
}
