package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
	"github.com/MrAinslay/fiber-rss-feed-cli/internal/utils"
	"golang.org/x/net/html"
)

func commandRead(cfg *ApiConfig, s string) error {
	rsp, err := cfg.ApiClient.HttpClient.Get(fmt.Sprintf("%s/posts/%s", cfg.ApiClient.BaseURL, s))
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(rsp.Body)
	params := api.Post{}
	if err := decoder.Decode(&params); err != nil {
		return err
	}

	defer rsp.Body.Close()

	if params.ErrorMsg != "" {
		return errors.New(params.ErrorMsg)
	}

	htmlRsp, err := cfg.ApiClient.HttpClient.Get(params.URL)
	if err != nil {
		return err
	}

	defer htmlRsp.Body.Close()

	body, err := html.Parse(htmlRsp.Body)
	if err != nil {
		return err
	}

	utils.ProcessArticleBody(body)

	return nil
}
