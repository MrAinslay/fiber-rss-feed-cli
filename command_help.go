package main

import (
	"fmt"
)

func commandHelp(cfg *ApiConfig, s string) error {
	fmt.Println("\nWelcome to Fiber RSS Feed \nUsage:")
	fmt.Println("")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
