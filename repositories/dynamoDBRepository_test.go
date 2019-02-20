package repositories

import (
	"testing"

	"github.com/bgiulianetti/api-mutantes/individual"
)

func TestAdd(T *testing.T) {

	mutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	mutant := individual.Individual{DNA: mutantDNA, ID: "123456"}

	service, _ := NewPersistenceService()
	service.Add(mutant, "mutant")

	mutantGotten, _ := service.Get(mutant.ID, "mutant")

	if mutantGotten.ID != mutant.ID {
		T.Error("El id debería ser el mismo")
	}
}

func TestGetCount(T *testing.T) {

	service, _ := NewPersistenceService()
	mutantCount, _ := service.GetCount("mutant")

	mutantDNA := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	mutant := individual.Individual{DNA: mutantDNA, ID: "123456"}

	service.Add(mutant, "mutant")

	mutantCountAfterAdd, _ := service.GetCount("mutant")

	if mutantCountAfterAdd.Count != mutantCount.Count+1 {
		T.Error("El count debería ser el mismo")
	}
}

func TestGetStats(T *testing.T) {

	service, _ := NewPersistenceService()
	mutantCount, _ := service.GetCount("mutant")
	humanCount, _ := service.GetCount("human")

	stats, _ := service.GetStats()

	if stats.CountHuman != float64(humanCount.Count) {
		T.Error("El count de humanos debería ser el mismo")
	}
	if stats.CountMutant != float64(mutantCount.Count) {
		T.Error("El count debería ser el mismo")
	}

}
