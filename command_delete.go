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

func commandDelete(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")

	switch splitString[0] {
	case "user":
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users", cfg.ApiClient.BaseURL), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Println(params.Message)
	case "feed":
		if len(splitString) < 2 {
			return errors.New("not enough arguments")
		}

		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/feeds/%s", cfg.ApiClient.BaseURL, splitString[1]), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Println(params.Message)
	case "feed_follow":
		if len(splitString) < 2 {
			return errors.New("not enough arguments")
		}

		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/feed-follows/%s", cfg.ApiClient.BaseURL, splitString[1]), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Printf(params.Message)
	case "post-like":
		if len(splitString) < 2 {
			return errors.New("not enough arguments")
		}

		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/post-likes/%s", cfg.ApiClient.BaseURL, splitString[1]), bytes.NewReader([]byte("")))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("ApiKey %s", cfg.ApiKey))

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
		if err != nil {
			return err
		}

		defer rsp.Body.Close()

		decoder := json.NewDecoder(rsp.Body)
		params := api.DeleteMsg{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErrorMsg != "" {
			return errors.New(params.ErrorMsg)
		}

		fmt.Println(params.Message)
	}
	return nil
}
