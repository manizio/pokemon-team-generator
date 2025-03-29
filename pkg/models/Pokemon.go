package models

import (
	"encoding/json"
	"fmt"
	"pokemon-team-generator/pkg/requests"
	"pokemon-team-generator/pkg/utils"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Pokemon struct {
	ID int `json:"id"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Sprite struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
}

func (p *Pokemon) FormatName() {
	p.Species.Name = strings.Replace(p.Species.Name, "-", " ", -1)
	p.Species.Name = cases.Title(language.English, cases.Compact).String(p.Species.Name)
}

func GetSpeciesID(url string) (int, error) {
	res, err := requests.MakeRequest("GET", url, nil)

	if err != nil {
		return -1, err
	}
	defer res.Body.Close()

	var species struct {
		ID int `json:"id"`
	}
	if err = json.NewDecoder(res.Body).Decode(&species); err != nil {
		return -1, err
	}

	return species.ID, nil

}
func GetRandomPokemon() (Pokemon, error) {
	number := utils.GenerateRandomNumber()
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", number)

	res, err := requests.MakeRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	var p Pokemon
	err = json.NewDecoder(res.Body).Decode(&p)

	if err != nil {
		return Pokemon{}, err
	}

	return p, nil
}

func GetPokemonByID(id int) (Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", id)

	res, err := requests.MakeRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	var p Pokemon
	err = json.NewDecoder(res.Body).Decode(&p)

	if err != nil {
		return Pokemon{}, err
	}

	return p, nil
}

