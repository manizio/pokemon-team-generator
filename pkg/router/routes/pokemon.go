package routes

import (
	"net/http"
	"pokemon-team-generator/pkg/controllers"
)
var pokemonRoutes = []Route {
	{
		URI: "/",
		Method: http.MethodGet,
		F: controllers.LoadHomePage,
	},
	{
		URI: "/team",
		Method: http.MethodGet,
		F: controllers.CreateTeam,
	},
}
