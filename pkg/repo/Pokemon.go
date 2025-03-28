package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokemon-team-generator/pkg/models"
	"pokemon-team-generator/pkg/utils"
)

type PokeAPI struct {
	url string
}

func NewPokeAPI(url string) *PokeAPI {
	return &PokeAPI{url}
}

func (api PokeAPI) GetPokemon() (models.Pokemon, error) {
	number := utils.GenerateRandomNumber()
	api.url = fmt.Sprintf("%v%v", api.url, number)

	req, err := http.NewRequest("GET", api.url, nil)

	if err != nil {
		return models.Pokemon{}, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return models.Pokemon{}, err
	}
	defer res.Body.Close()

	var p models.Pokemon
	err = json.NewDecoder(res.Body).Decode(&p)

	if err != nil {
		return models.Pokemon{}, err
	}

	return p, nil
}

func (api PokeAPI) GetPokemonByID(id int) (models.Pokemon, error) {
	api.url = fmt.Sprintf("%v%v", api.url, id)

	req, err := http.NewRequest("GET", api.url, nil)

	if err != nil {
		return models.Pokemon{}, err
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return models.Pokemon{}, err
	}
	defer res.Body.Close()

	var p models.Pokemon
	err = json.NewDecoder(res.Body).Decode(&p)

	if err != nil {
		return models.Pokemon{}, err
	}

	return p, nil
}
