package main

import (
	"log"
	"net/http"
	"pokemon-team-generator/pkg/router"
)

func main() {

	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":3000", r))
}
