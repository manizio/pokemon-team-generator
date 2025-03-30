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

type SpeciesInfo struct {
	ID          int  `json:"id"`
	IsMythical  bool `json:"is_mythical"`
	IsLegendary bool `json:"is_legendary"`
}

type Pokemon struct {
	ID      int `json:"id"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Sprite struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	IsMythical  bool
	IsLegendary bool
	Types []PokemonType `json:"types"`
}

func (p *Pokemon) Format() {
	p.FormatName()
	p.FormatType()
}

func (p *Pokemon) FormatName() {
	p.Species.Name = strings.Replace(p.Species.Name, "-", " ", -1)
	p.Species.Name = cases.Title(language.English, cases.Compact).String(p.Species.Name)
}

func (p *Pokemon) FormatType() {
	for _, t := range p.Types {
		t.Type.Name = cases.Title(language.English, cases.Compact).String(t.Type.Name)
	}
}

func GetSpeciesInfo(url string) (SpeciesInfo, error) {
	res, err := requests.MakeRequest("GET", url, nil)

	if err != nil {
		return SpeciesInfo{}, err
	}
	defer res.Body.Close()

	var species SpeciesInfo

	if err = json.NewDecoder(res.Body).Decode(&species); err != nil {
		return SpeciesInfo{}, err
	}

	return species, nil

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
