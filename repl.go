package main

import (
	"bufio"
	"os"
	"strings"
	"github.com/fatih/color"
	"github.com/kalim-Asim/pokedexcli/internal/commands"
	"github.com/kalim-Asim/pokedexcli/internal/pokeapi"
)

func startREPL(cfg *pokeapi.Config) {
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
		availableCommands := commands.GetCommands()
		actualCommand, ok := availableCommands[userCommand]
		args := words[1:]

		if !ok {
			color.Red("Unknown commmand !!")
			continue
		}

		err := actualCommand.Callback(cfg, args)
		if err != nil {
			color.Red("Error occured", err)
		}
	}
}

func cleanInput(s string) []string {
	str := strings.TrimSpace(s)
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
