package main

import (
	"log"
	"net/http"

	"github.com/bgiulianetti/api-mutantes/api"
)

func main() {

	router := api.NewRouter()
	server := http.ListenAndServe(":5000", router)
	log.Fatal(server)
}
