package main

import "fmt"

func CommandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.pokedex {
		fmt.Println(" - " + p.Name)
	}
	return nil
}
