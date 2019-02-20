package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
)

//ConvertToUpperCase transformo un array de string a uppercase
func ConvertToUpperCase(array []string) []string {
	for i := 0; i < len(array); i++ {
		array[i] = strings.ToUpper(array[i])
	}
	return array
}

//GenerateTimeStamp Genera un timestamp para usar como id en la bd
func GenerateTimeStamp() string {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid

}

// ConcatenateStringArray ...
func ConcatenateStringArray(array []string) string {
	cadena := ""
	for _, element := range array {
		cadena += element
	}
	return cadena
}
