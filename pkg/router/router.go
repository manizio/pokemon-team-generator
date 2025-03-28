package router 

import (
	"pokemon-team-generator/pkg/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}

