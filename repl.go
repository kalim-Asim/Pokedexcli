package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error 
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		userCommand := words[0]
		commands := getCommands()
		actualCommand, ok := commands[userCommand]
		if !ok {
			fmt.Println("Unknown command")
			continue 
		}

		actualCommand.callback()
	}
}

func cleanInput(s string) []string {
	str := strings.TrimSpace(s)
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: callbackExit,
		},
	}
}
