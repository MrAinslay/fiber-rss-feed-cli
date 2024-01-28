package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
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

	body, err := io.ReadAll(htmlRsp.Body)
	if err != nil {
		return err
	}

	log.Println(string(body))
	parseAndShow(string(body))

	return nil
}

func parseAndShow(text string) {

	tkn := html.NewTokenizer(strings.NewReader(text))

	var isTd bool
	var n int

	for {

		tt := tkn.Next()

		switch {

		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:

			t := tkn.Token()
			isTd = t.Data == "td"

		case tt == html.TextToken:

			t := tkn.Token()

			if isTd {

				fmt.Printf("%s ", t.Data)
				n++
			}

			if isTd && n%3 == 0 {

				fmt.Println()
			}

			isTd = false
		}
	}
}
