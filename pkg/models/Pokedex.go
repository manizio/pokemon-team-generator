package models

import (
	"encoding/json"
	"fmt"
	"pokemon-team-generator/pkg/requests"
)

type Pokedex struct {
	ID             int `json:"id"`
	PokemonEntries []struct {
		PokemonSpecies struct {
			Url string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"pokemon_entries"`
}

func GetPokedex(id int) (Pokedex, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokedex/%d", id)

	res, err := requests.MakeRequest("GET", url, nil)

	if err != nil {
		return Pokedex{}, err
	}

	defer res.Body.Close()

	var p Pokedex
	if err = json.NewDecoder(res.Body).Decode(&p); err != nil {
		return Pokedex{}, err
	}

	return p, nil
}
