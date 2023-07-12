package main

import (
	"encoding/json"
	"net/http"

	"github.com/braydend/pokecompare/api"
)

const PORT = ":8000"

type RandomPokemonResponse struct {
	Name   string
	ID     int
	Sprite string
	Types  []string
}

func main() {
	http.HandleFunc("/random-pokemon", func(w http.ResponseWriter, r *http.Request) {
		pokemon, err := api.GetRandomPokemon()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}

		var types []string

		for _, pokemonType := range pokemon.Types {
			types = append(types, pokemonType.Type.Name)
		}

		resp, err := json.Marshal(RandomPokemonResponse{
			Name:   pokemon.Name,
			ID:     pokemon.ID,
			Sprite: pokemon.Sprites.FrontDefault,
			Types:  types,
		})

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		w.Write(resp)
	})

	http.ListenAndServe(PORT, nil)
}
