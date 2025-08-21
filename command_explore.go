package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	area := args[0]
	explore, err := cfg.pokeapiClient.ExploreLocation(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	for _, p := range explore.PokemonEncounters {
		fmt.Println("- " + p.Pokemon.Name)
	}

	return nil
}
