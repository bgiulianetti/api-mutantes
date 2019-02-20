package main

import (
	"api"
	"log"
	"net/http"
)

func main() {

	router := api.NewRouter()
	server := http.ListenAndServe(":5000", router)
	log.Fatal(server)
}
