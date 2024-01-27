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

func commandUpdate(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")
	switch splitString[0] {
	case "user":
		jsonBody := fmt.Sprintf(`{"name": "%s", "password": "%s"}`, splitString[1], splitString[2])
		bodyReader := bytes.NewReader([]byte(jsonBody))
		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users", cfg.ApiClient.BaseURL), bodyReader)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", cfg.ApiKey)

		rsp, err := cfg.ApiClient.HttpClient.Do(req)
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

		fmt.Printf("\n\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	}
	return nil
}
