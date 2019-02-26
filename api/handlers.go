package api

import (
	"encoding/json"
	"net/http"

	"github.com/bgiulianetti/api-mutantes/utils"

	"github.com/bgiulianetti/api-mutantes/individual"
	"github.com/bgiulianetti/api-mutantes/repositories"
)

//ApiError ..
type ApiError struct {
	Message string `json:"message"`
}

//Stats me dice la canatidad de adn validos y no validos
func Stats(w http.ResponseWriter, r *http.Request) {

	service, err := repositories.NewPersistenceService()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
		_ = repositories.UploadToS3(individual.Individual{}, err.Error())
		return
	}

	stats, err := service.GetStats()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
		_ = repositories.UploadToS3(individual.Individual{}, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(stats)
}

//DetectMutant detecta si el adn pasado es valido o no
func DetectMutant(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var individual individual.Individual
	err := decoder.Decode(&individual)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
		_ = repositories.UploadToS3(individual, err.Error())
		return
	}

	individual.DNA = utils.ConvertToUpperCase(individual.DNA)
	if !individual.IsDnaFormatValid() {
		w.WriteHeader(400)
		return
	}

	service, err := repositories.NewPersistenceService()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
		_ = repositories.UploadToS3(individual, err.Error())
		return
	}

	if individual.IsMutant() {

		err = service.Add(individual, "mutant")
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
			_ = repositories.UploadToS3(individual, err.Error())
			return
		}
		w.WriteHeader(200)
	} else {

		err = service.Add(individual, "human")
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(ApiError{Message: err.Error()})
			_ = repositories.UploadToS3(individual, err.Error())
			return
		}
		w.WriteHeader(403)
	}
}

// Health check...
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
