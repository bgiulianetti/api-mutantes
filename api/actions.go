package api

import (
	"encoding/json"
	"net/http"
)

//Stats me dice la canatidad de adn validos y no validos
func Stats(w http.ResponseWriter, r *http.Request) {
	stats, err := GetIndividualStats()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

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
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	if isMutant(individual.DNA) {
		err = AddIndividual(individual.DNA, "mutant")
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(200)
	} else {
		err = AddIndividual(individual.DNA, "human")
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(403)
	}
}

// Health ...
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
