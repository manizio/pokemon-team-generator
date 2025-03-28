package controllers

import (
	"net/http"
	"pokemon-team-generator/pkg/models"
	"pokemon-team-generator/pkg/repo"
	"pokemon-team-generator/pkg/utils"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team []models.Pokemon

	url := "https://pokeapi.co/api/v2/pokemon/"
	repo := repo.NewPokeAPI(url)
	for range 6 {
		pokemon, err := repo.GetPokemon()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		pokemon.FormatName()
		team = append(team, pokemon)
	}

	utils.ExecTemplate(w, "home", team)

}
