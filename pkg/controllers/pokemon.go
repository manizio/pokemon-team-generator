package controllers

import (
	"net/http"
	"pokemon-team-generator/pkg/models"
	"pokemon-team-generator/pkg/utils"
)

// Cria um time de 6 Pokémons Aleatórios
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	teamChannel := make(chan models.Pokemon, 6)
	var team []models.Pokemon

	for range 6 {
		go func() {
			pokemon, err := models.GetRandomPokemon()
			if err != nil {
				return
			}
			pokemon.FormatName()

			teamChannel <- pokemon
		}()
	}

	for range 6 {
		team = append(team, <-teamChannel)
	}
	utils.ExecTemplate(w, "home", team)
}
