package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI string
	Method string
	F func(http.ResponseWriter, *http.Request)
}

func Configure(r *mux.Router) *mux.Router {
	routes := pokemonRoutes

	for _, route := range routes {
		r.HandleFunc(
			route.URI,
			route.F,
		).Methods(route.Method)
	}

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))

	return r
}
