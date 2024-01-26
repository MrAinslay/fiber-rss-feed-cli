package main

import (
	"errors"
	"strings"
)

func commandSet(cfg *ApiConfig, s string) error {
	splitString := strings.Split(s, " ")
	switch splitString[1] {
	case "apikey":
		if splitString[2] == "" {
			return errors.New("invalid api key")
		}
		cfg.ApiKey = splitString[2]
	}
	return nil
}
