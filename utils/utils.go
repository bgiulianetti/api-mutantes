package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bgiulianetti/api-mutantes/individual"
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
	_, _ = rand.Read(b)

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

// LogDNA ...
func LogDNA(individualToPersist individual.Individual, reason string) {

	_, _ = s3.New(session.Must(session.NewSession(&aws.Config{Region: aws.String("sa-east-1")}))).PutObject(&s3.PutObjectInput{
		Bucket: aws.String("api-mutantes-failed-request"),
		Key:    aws.String(GenerateTimeStamp() + ".txt"),
		Body:   bytes.NewReader([]byte(time.Now().Format(time.RFC3339) + "  - ID: " + individualToPersist.ID + " - " + reason)),
	})
}
