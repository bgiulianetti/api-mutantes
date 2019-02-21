package utils

import "testing"

func TestGenerateTimeStamp(T *testing.T) {

	timeStampA := GenerateTimeStamp()
	timeStampB := GenerateTimeStamp()

	if timeStampA == timeStampB {
		T.Error("TimeStamps deberían ser diferences")
	}

}

func TestConvertToUpperCase(T *testing.T) {

	arrayLowerCase := []string{
		"abc",
		"def"}

	arrayUpperCaseToCompare := []string{
		"ABC",
		"DEF"}

	arrayUpperCase := ConvertToUpperCase(arrayLowerCase)

	for index, element := range arrayUpperCase {
		if element != arrayUpperCaseToCompare[index] {
			T.Error("Las cadenas deberían ser iguales")
		}
	}
}

func TestConcatenateDNA(T *testing.T) {

	array := []string{
		"abc",
		"def"}

	concatenatedArray := ConcatenateStringArray(array)

	if concatenatedArray != "abcdef" {
		T.Error("Las cadenas deberían ser iguales")
	}
}
