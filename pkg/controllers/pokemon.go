package controllers

import (
	"net/http"
	"pokemon-team-generator/pkg/models"
	"pokemon-team-generator/pkg/utils"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team []models.Pokemon

	for range 6 {
		pokemon, err := models.GetRandomPokemon()
		if err != nil {
			http.Error(w, "home", http.StatusInternalServerError)
		}
		pokemon.FormatName()
		team = append(team, pokemon)
	}
	utils.ExecTemplate(w, "home", team)

}
