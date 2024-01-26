package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

		decoder := json.NewDecoder(rsp.Body)
		params := api.User{}
		if err := decoder.Decode(&params); err != nil {
			return err
		}

		if params.ErroMsg != "" {
			return errors.New(params.ErroMsg)
		}

		log.Printf("\n\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	case "feed":
		jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "url": "%s"}`, splitString[1], splitString[2]))
		bodyReader := bytes.NewReader(jsonBody)
		rsp, err := cfg.ApiClient.HttpClient.Post(fmt.Sprintf("%s/feeds", cfg.ApiClient.BaseURL), "application/json", bodyReader)
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
	}

	return nil
}
