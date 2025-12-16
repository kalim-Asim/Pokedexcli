package main

import (
	"os"
	"github.com/fatih/color"
)

func callbackExit(cfg* config) error {
	color.Red("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}