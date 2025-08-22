package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func CommandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", name)
	chanceToCatch := rand.Intn(pokemonResp.BaseExperience)
	if chanceToCatch < 40 {
		fmt.Println(pokemonResp.Name + " escaped!")
		return nil
	}

	fmt.Println(pokemonResp.Name + " was caught!")
	cfg.pokedex[pokemonResp.Name] = pokemonResp

	return nil
}
