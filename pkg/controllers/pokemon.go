package controllers

import (
	"fmt"
	"net/http"
	"pokemon-team-generator/pkg/components"
	"pokemon-team-generator/pkg/models"
	"pokemon-team-generator/pkg/utils"
	"strconv"
)

// Carrega a página inicial
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	components.HomePage().Render(r.Context(), w)
}

// Cria um time de 6 Pokémons aleatórios
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	pokedexID, err := strconv.Atoi(r.FormValue("pokedex"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fullPokedex, err := getFullPokedex(pokedexID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	teamChannel := make(chan models.Pokemon, 6)
	var team []models.Pokemon
	for range 6 {
		dexLen := len(fullPokedex.PokemonEntries)
		randomIdx := utils.GenerateRandomNumberN(dexLen)
		pokeInfo, _ := models.GetSpeciesInfo(
			fullPokedex.PokemonEntries[randomIdx].PokemonSpecies.Url,
		)

		go func(pokeInfo models.SpeciesInfo) {
			pokemon, err := models.GetPokemonByID(pokeInfo.ID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			pokemon.Format()

			pokemon.IsMythical = pokeInfo.IsMythical
			pokemon.IsLegendary = pokeInfo.IsLegendary

			teamChannel <- pokemon
		}(pokeInfo)
	}

	for range 6 {
		team = append(team, <-teamChannel)
	}

	components.Team(team).Render(r.Context(), w)
}

// Troca o Pokémon selecionado
func SwapPokemon(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	pokedexID, err := strconv.Atoi(r.FormValue("pokedex"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fullPokedex, err := getFullPokedex(pokedexID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dexLen := len(fullPokedex.PokemonEntries)
	randomIdx := utils.GenerateRandomNumberN(dexLen)
	pokeInfo, _ := models.GetSpeciesInfo(
		fullPokedex.PokemonEntries[randomIdx].PokemonSpecies.Url,
	)

	pokemon, err := models.GetPokemonByID(pokeInfo.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	pokemon.Format()

	pokemon.IsMythical = pokeInfo.IsMythical
	pokemon.IsLegendary = pokeInfo.IsLegendary

	components.PokemonDiv(pokemon).Render(r.Context(), w)

}

// Obtém a Pokédex do jogo selecionado
func getFullPokedex(pokedexID int) (models.Pokedex, error) {
	var fullPokedex models.Pokedex
	var lastID int

	switch pokedexID {
	case 16:
		lastID = 20
	case 21:
		lastID = 25
	default:
		lastID = pokedexID
	}

	for i := pokedexID; i <= lastID; i++ {
		pokedex, err := models.GetPokedex(pokedexID)
		if err != nil {
			fmt.Println("error: ", err)
			return models.Pokedex{}, err
		}
		if fullPokedex.PokemonEntries == nil {
			fullPokedex = pokedex
		} else {
			fullPokedex.PokemonEntries = append(fullPokedex.PokemonEntries, pokedex.PokemonEntries...)
		}
	}

	return fullPokedex, nil
}
