package main

import (
	"fmt"
	"log"

	"github.com/braydend/pokecompare/api"
)

func main() {
	pokemon, err := api.GetRandomPokemon()

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to get a pokemon %s", err))
	}

	fmt.Printf("Successfully got %s", pokemon.Name)
}
