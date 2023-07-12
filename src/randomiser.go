package src

import (
	"math/rand"

	"github.com/braydend/pokecompare/api"
)

func generateRandomMatchup() (api.PokemonResponse, api.PokemonResponse, error) {
	randomPokemonOne, err := api.GetRandomPokemon()

	if err != nil {
		return api.PokemonResponse{}, api.PokemonResponse{}, err
	}

	randomPokemonTwo, err := api.GetRandomPokemon()

	if err != nil {
		return api.PokemonResponse{}, api.PokemonResponse{}, err
	}

	return randomPokemonOne, randomPokemonTwo, nil
}

func GenerateMatchup(votedIds []int) (api.PokemonResponse, api.PokemonResponse, error) {
	if len(votedIds) == 0 {
		return generateRandomMatchup()
	}

	randomIndex := rand.Intn(len(votedIds))
	previouslySelectedId := votedIds[randomIndex]

	previousPokemon, err := api.GetPokemonById(previouslySelectedId)

	if err != nil {
		return api.PokemonResponse{}, api.PokemonResponse{}, err
	}

	randomPokemon, err := api.GetRandomPokemon()

	if err != nil {
		return api.PokemonResponse{}, api.PokemonResponse{}, err
	}

	return previousPokemon, randomPokemon, nil
}
