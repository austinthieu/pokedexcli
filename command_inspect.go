package main

import (
	"errors"
	"fmt"
)

func CommandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("enter a valid pokemon name")
	}

	name := args[0]

	pokemon, exists := cfg.pokedex[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" - " + t.Type.Name)
	}
	return nil
}
