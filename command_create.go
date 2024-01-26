package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
)

func commandCreate(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")
	if splitString[0] == "user" {
		jsonBody := []byte(fmt.Sprintf(`{"name": "%s", "password": "%s"}`, splitString[1], splitString[2]))
		bodyReader := bytes.NewReader(jsonBody)
		rsp, err := cfg.ApiClient.HttpClient.Post(fmt.Sprintf("%s/users", cfg.ApiClient.BaseURL), "application/json", bodyReader)
		if err != nil {
			log.Println(err)
		}

		decoder := json.NewDecoder(rsp.Body)
		params := api.User{}
		if err := decoder.Decode(&params); err != nil {
			log.Println(err)
		}

		if params.ErroMsg != "" {
			log.Println(params.ErroMsg)
			return nil
		}

		log.Printf("\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n", params.ID, params.CreatedAt, params.Name, params.APIKey)
	}

	return nil
}
