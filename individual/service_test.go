package individual

import "testing"

func TestIsMutantWithMutantDNA(T *testing.T) {

	mutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	mutant := Individual{DNA: mutantDNA, ID: "123456"}

	if !mutant.IsMutant() {
		T.Error("El adn debería ser mutante")
	}
}

func TestIsMutantWithMutantDNAHorizontal(T *testing.T) {

	mutantDNA := []string{
		"AAAAAA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}
	mutant := Individual{DNA: mutantDNA, ID: "123456"}

	if !mutant.IsMutant() {
		T.Error("El adn debería ser mutante")
	}
}

func TestIsMutantWithMutantDNASlash(T *testing.T) {

	mutantDNA := []string{
		"AATAAA",
		"CAGTAC",
		"TTTAGT",
		"AGAAGG",
		"CCGCGA",
		"TCACGG"}
	mutant := Individual{DNA: mutantDNA, ID: "123456"}

	if !mutant.IsMutant() {
		T.Error("El adn debería ser mutante")
	}
}

func TestIsMutantWithHumanDNA(T *testing.T) {
	humanDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGACCG",
		"TACCTA",
		"TCACTG"}

	human := Individual{DNA: humanDNA, ID: "123456"}
	if human.IsMutant() {
		T.Error("El adn debería ser humano")
	}
}

func TestIsDnaFormatValidWithValidDNA(T *testing.T) {
	validDNA := []string{
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	individualWithValidDNA := Individual{DNA: validDNA, ID: "123456"}
	if !individualWithValidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser valido")
	}
}

func TestIsDnaFormatValidWithFiveCharString(T *testing.T) {
	invalidDNA := []string{
		"AAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	individualWithInvalidDNA := Individual{DNA: invalidDNA, ID: "123456"}
	if individualWithInvalidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithFiveStringLength(T *testing.T) {
	invalidDNA := []string{
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	individualWithInvalidDNA := Individual{DNA: invalidDNA, ID: "123456"}
	if individualWithInvalidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithWrongCharacter(T *testing.T) {
	invalidDNA := []string{
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAX"}

	individualWithInvalidDNA := Individual{DNA: invalidDNA, ID: "123456"}
	if individualWithInvalidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithEmptyArray(T *testing.T) {
	invalidDNA := []string{}

	individualWithInvalidDNA := Individual{DNA: invalidDNA, ID: "123456"}
	if individualWithInvalidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithWrongChars(T *testing.T) {
	invalidDNA := []string{
		"AAXAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	individualWithInvalidDNA := Individual{DNA: invalidDNA, ID: "123456"}
	if individualWithInvalidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithWrongCharsTwo(T *testing.T) {
	invalidDNA := []string{
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAQ"}

	individualWithInvalidDNA := Individual{DNA: invalidDNA, ID: "123456"}
	if individualWithInvalidDNA.IsDnaFormatValid() {
		T.Error("El adn debería ser invalido")
	}
}
