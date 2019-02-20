package models

import (
	"testing"
)

func TestIsMutantWithMutantDNA(T *testing.T) {
	mutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG"}

	if !isMutant(mutantDNA) {
		T.Error("El adn debería ser mutante")
	}
}

func TestIsMutantWithHumanDNA(T *testing.T) {
	mutantDNA := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGACCG",
		"TACCTA",
		"TCACTG"}

	if isMutant(mutantDNA) {
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

	if !isDnaFormatValid(validDNA) {
		T.Error("El adn debería ser valido")
	}
}

func TestIsDnaFormatValidWithFiveCharString(T *testing.T) {
	validDNA := []string{
		"AAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	if isDnaFormatValid(validDNA) {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithFiveStringLength(T *testing.T) {
	validDNA := []string{
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	if isDnaFormatValid(validDNA) {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithEmptyArray(T *testing.T) {
	validDNA := []string{}

	if isDnaFormatValid(validDNA) {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithWrongChars(T *testing.T) {
	validDNA := []string{
		"AAXAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA"}

	if isDnaFormatValid(validDNA) {
		T.Error("El adn debería ser invalido")
	}
}

func TestIsDnaFormatValidWithWrongCharsTwo(T *testing.T) {
	validDNA := []string{
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAA",
		"AAAAAQ"}

	if isDnaFormatValid(validDNA) {
		T.Error("El adn debería ser invalido")
	}
}
