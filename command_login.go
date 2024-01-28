package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
)

func commandLogin(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")

	if len(splitString) < 2 {
		return errors.New("not enough arguments")
	}

	jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "password": "%s"}`, splitString[0], splitString[1]))
	bodyReader := bytes.NewReader(jsonBody)
	rsp, err := cfg.ApiClient.HttpClient.Post(fmt.Sprintf("%s/login", cfg.ApiClient.BaseURL), "application/json", bodyReader)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(rsp.Body)
	params := api.User{}
	if err := decoder.Decode(&params); err != nil {
		return err
	}

	defer rsp.Body.Close()

	if params.ErrorMsg != "" {
		return errors.New(params.ErrorMsg)
	}

	cfg.ApiKey = params.ApiKey
	fmt.Printf("\nSet API key to %s\n", params.ApiKey)
	fmt.Printf("\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	return nil
}
