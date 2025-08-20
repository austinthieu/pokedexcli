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

var commands map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func init() {
	commands = map[string]cliCommand{
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
	}
}

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
		command := cleanedInput[0]

		cmd, exists := commands[command]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Command '%s' failed: %v\n", command, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowercaseString := strings.ToLower(text)
	return strings.Fields(lowercaseString)
}
