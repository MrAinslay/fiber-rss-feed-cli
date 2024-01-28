package main

import (
	"errors"
	"fmt"
	"strings"
)

func commandSet(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")
	if len(splitString) < 2 {
		cfg.ApiKey = ""
		fmt.Println("Set API key to <EMPTY>")
		return nil
	}
	switch splitString[0] {
	case "apikey":

		if splitString[1] == "" {
			return errors.New("invalid api key")
		}
		cfg.ApiKey = splitString[1]
		fmt.Printf("Set API key to %s", cfg.ApiKey)
	}
	return nil
}
