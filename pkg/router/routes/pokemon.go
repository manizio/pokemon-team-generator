package routes

import (
	"net/http"
	"pokemon-team-generator/pkg/controllers"
)
var pokemonRoutes = []Route {
	{
		URI: "/team",
		Method: http.MethodGet,
		F: controllers.CreateTeam,
	},
}
