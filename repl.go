package main

import (
	"bufio"
	"os"
	"strings"
	"github.com/fatih/color"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	color.Green("Welcome to Pokemon World!")
	
	for {
		cyan := color.New(color.FgCyan).PrintfFunc()
		cyan("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		userCommand := words[0]
		availableCommands := getCommands()
		actualCommand, ok := availableCommands[userCommand]
		args := words[1:]

		if !ok {
			color.Red("Unknown commmand !!")
			continue
		}

		err := actualCommand.callback(cfg, args)
		if err != nil {
			color.Red("Error occured", err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists next 20 location",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous 20 location",
			callback:    callbackMapb,
		},
		"explore": {
			name:					"explore",
			description: 	"Explore any particular location",
			callback: 		callbackExplore,
		},
		"catch": {
			name: 			"catch",
			description: "Catch a particular pokemon",
			callback:    callbackCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(s string) []string {
	str := strings.TrimSpace(s)
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
