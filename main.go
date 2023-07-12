package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/braydend/pokecompare/src"
)

type pokemonData struct {
	Name   string
	ID     int
	Sprite string
	Types  []string
}

type MatchupResponse struct {
	PokemonOne pokemonData `json:"pokemonOne`
	PokemonTwo pokemonData `json:"pokemonTwo`
}

type RequestBody struct {
	PreviousIds []int `json:"previousIds"`
}

func main() {
	PORT := ":" + os.Getenv("PORT")

	if PORT == ":" {
		PORT = ":8000"
	}

	fs := http.FileServer(http.Dir("./web/dist"))

	http.HandleFunc("/generate-matchup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			w.Write([]byte("POST only!"))
			return
		}

		var requestBody RequestBody

		requestBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Failed to read body." + err.Error()))
			return
		}

		err = json.Unmarshal(requestBytes, &requestBody)

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Failed to parse body." + err.Error()))
			return
		}

		pokemonOne, pokemonTwo, err := src.GenerateMatchup(requestBody.PreviousIds)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		var pokemonOneTypes []string
		var pokemonTwoTypes []string

		for _, pokemonType := range pokemonOne.Types {
			pokemonOneTypes = append(pokemonOneTypes, pokemonType.Type.Name)
		}

		for _, pokemonType := range pokemonTwo.Types {
			pokemonTwoTypes = append(pokemonTwoTypes, pokemonType.Type.Name)
		}

		resp, err := json.Marshal(MatchupResponse{
			PokemonOne: pokemonData{
				Name:   pokemonOne.Name,
				ID:     pokemonOne.ID,
				Sprite: pokemonOne.Sprites.FrontDefault,
				Types:  pokemonOneTypes,
			},
			PokemonTwo: pokemonData{
				Name:   pokemonTwo.Name,
				ID:     pokemonTwo.ID,
				Sprite: pokemonTwo.Sprites.FrontDefault,
				Types:  pokemonTwoTypes,
			}})

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(resp)
	})

	http.Handle("/", fs)

	log.Default().Println("Running on port " + PORT)
	http.ListenAndServe(PORT, nil)
}
