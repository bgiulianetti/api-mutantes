package individual

import "testing"

func TestGetDNAFromindividual(T *testing.T) {

	mutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	mutant := Individual{DNA: mutantDNA, ID: "123456"}

	for index, element := range mutantDNA {
		if element != mutant.DNA[index] {
			T.Error("Los pares de adn deberían ser iguales")
		}
	}
}
