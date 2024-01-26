package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MrAinslay/fiber-rss-feed-cli/internal/api"
)

type ApiConfig struct {
	ApiClient api.Client
	ApiKey    string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *ApiConfig, s string) error
}

func cleanInput(s string) string {
	output := strings.TrimSpace(s)
	output = strings.ToLower(output)
	return output
}

func printPrompt() {
	fmt.Print("Rss-Feed > ")
}

func startRepl(cfg *ApiConfig) {
	commands := getCommands()

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := cleanInput(reader.Text())
		splitText := strings.Split(text, " ")
		if command, exists := commands[splitText[0]]; exists {
			if len(splitText) == 4 {
				err := command.callback(cfg, fmt.Sprintf("%s %s %s", splitText[1], splitText[2], splitText[3]))
				if err != nil {
					log.Println(err)
				}
			} else if len(splitText) == 3 {
				err := command.callback(cfg, fmt.Sprintf("%s %s", splitText[1], splitText[2]))
				if err != nil {
					log.Println(err)
				}
			} else if len(splitText) > 1 {
				err := command.callback(cfg, splitText[1])
				if err != nil {
					log.Println(err)
				}

			} else {
				err := command.callback(cfg, splitText[0])
				if err != nil {
					log.Println(err)
				}
			}
		} else {
			log.Println("Unknown Command")
		}
		printPrompt()
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Closes the CLI",
			callback:    commandExit,
		},
		"create": {
			name:        "create",
			description: "Creates either a user or a feed depending on the selected option",
			callback:    commandCreate,
		},
		"login": {
			name:        "login",
			description: "Login as a user and sets the api key\n(usage: login <NAME> <PASSWORD>)",
			callback:    commandLogin,
		},
		"set": {
			name:        "set",
			description: "Set API Key for commands that need authorization\n(usage: set apikey <API_KEY>)",
			callback:    commandSet,
		},
		"get": {
			name:        "get",
			description: "Get feeds, posts, feed follows, post likes or user information",
			callback:    commandGet,
		},
		"update": {
			name:        "update",
			description: "Update information of a user\n(usage: update user <NAME> <PASSWORD)",
			callback:    commandUpdate,
		},
	}
}
