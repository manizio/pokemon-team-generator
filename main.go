package main

import (
	"log"
	"net/http"
	"pokemon-team-generator/pkg/middleware"
	"pokemon-team-generator/pkg/router"
)

func main() {

	r := router.GenerateRouter()
	r.Use(middleware.Logger)
	log.Fatal(http.ListenAndServe(":3000", r))
}
