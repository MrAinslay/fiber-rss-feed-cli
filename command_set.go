package main

import "errors"

func commandSet(cfg *ApiConfig, s string) error {
	if s == "" {
		return errors.New("invalid api key")
	}
	cfg.ApiKey = s
	return nil
}
