package main

import (
	"log"
	"net/http"
	"pokemon-team-generator/pkg/router"
	"pokemon-team-generator/pkg/utils"
)

func main() {

	utils.LoadTemplates()
	r := router.GenerateRouter()

	log.Fatal(http.ListenAndServe(":3000", r))
}
