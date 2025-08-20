package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Make API GET requests to the PokeAPI

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		cleanedInput := cleanInput(line)
		commandName := cleanedInput[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowercaseString := strings.ToLower(text)
	return strings.Fields(lowercaseString)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

type config struct {
	NextURL     string
	previousURL string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
	}
}
