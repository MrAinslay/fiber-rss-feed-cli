package main

import "log"

func commandHelp(cfg *ApiConfig, s string) error {
	log.Println("\nWelcome to Fiber RSS Feed \nUsage:")
	log.Println("")
	for _, cmd := range getCommands() {
		log.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	log.Println("")
	return nil
}
