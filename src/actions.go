package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Stats me dice la canatidad de adn validos y no validos
func Stats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Las estadisticas todavia no estan guardadas")
}

//DetectMutant detecta si el adn pasado es valido o no
func DetectMutant(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var individual Individual
	err := decoder.Decode(&individual)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if isMutant(individual.DNA) {
		AddMutant(individual.DNA)
		w.WriteHeader(200)
	} else {
		AddHuman()
		w.WriteHeader(403)
	}
}
