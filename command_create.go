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
		var jsonErr string
		if err := decoder.Decode(&jsonErr); err == nil {
			log.Println(jsonErr)
			return nil
		}
		if err := decoder.Decode(&params); err != nil {
			log.Println(err)
		}

		log.Printf("\nID: %s\nCreated At: %s\nName: %s\nApi Key: %s\n", params.Id, params.CreatedAt, params.Name, params.ApiKey)
	}

	return nil
}
