package api

import (
	"encoding/json"
	"net/http"
)

//Stats me dice la canatidad de adn validos y no validos
func Stats(w http.ResponseWriter, r *http.Request) {
	stats := GetIndividualStats()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(stats)
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
		AddIndividual(individual.DNA, "mutant")
		w.WriteHeader(200)
	} else {
		AddIndividual(individual.DNA, "human")
		w.WriteHeader(403)
	}
}

// Health ...
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
