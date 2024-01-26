package main

import "os"

func commandExit(cfg *ApiConfig, s string) error {
	os.Exit(0)
	return nil
}
