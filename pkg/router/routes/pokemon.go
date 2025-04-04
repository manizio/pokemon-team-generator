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
		Method: http.MethodPost,
		F: controllers.CreateTeam,
	},
	{
		URI: "/swap",
		Method: http.MethodPost,
		F: controllers.SwapPokemon,
	},
	{
		URI: "/lock",
		Method: http.MethodPost,
		F: controllers.LockPokemon,
	},
	{
		URI: "/unlock",
		Method: http.MethodPost,
		F: controllers.UnlockPokemon,
	},
}
