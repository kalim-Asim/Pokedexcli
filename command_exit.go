package main

import (
	"os"
	"fmt"
)

func callbackExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}