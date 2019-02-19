package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// CreateSeason crea una sesion de conexión a dynamodb
func CreateSeason() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String("sa-east-1")})
	svc := dynamodb.New(sess)
	return svc
}

// CheckifTableExists chequea si una tabla exite en dynamoDB
func CheckifTableExists(svc *dynamodb.DynamoDB) {
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("mutant"),
	}
	result, err := svc.DescribeTable(req)
	if err != nil {
		fmt.Printf("%s", err)
	}
	table := result.Table
	fmt.Printf("%+v\n", table)
}

// AddMutant agrega un mutante a dynamodb
func AddMutant(dna []string) {

	mutant := Individual{dna, GenerateTimeStamp()}
	item, err := dynamodbattribute.MarshalMap(mutant)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("mutant"),
	}
	svc := CreateSeason()
	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Error al llamar guardar en la BD")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Agregast el elemento con exito")
}

// AddHuman agrega un humano a un contador local
func AddHuman() {

}
