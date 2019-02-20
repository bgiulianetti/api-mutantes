package api

import (
	"fmt"

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

// AddIndividual agrega un mutante a dynamodb
func AddIndividual(dna []string, individualType string) error {

	individual := Individual{dna, GenerateTimeStamp()}
	item, err := dynamodbattribute.MarshalMap(individual)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(individualType),
	}
	svc := CreateSeason()
	_, err = svc.PutItem(input)

	if err != nil {
		return err
	}
	err = IncrementIndividualCount(svc, individualType)
	if err != nil {
		return err
	}
	return nil
}

// GetIndividualStats ...
func GetIndividualStats() IndividualStats {

	svc := CreateSeason()
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String("individualCount"),
	})

	if err != nil {
		fmt.Println("failed to make Query API call", err)
	}

	items := []IndividualCount{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		fmt.Println("failed to unmarshal Query result items", err)
	}
	response := IndividualStats{}

	for _, item := range items {
		if item.ID == "human" {
			response.CountHuman = item.Count
		}
		if item.ID == "mutant" {
			response.CountMutant = item.Count
		}
	}
	response.Ratio = float32(response.CountMutant / response.CountHuman)
	return response
}

// IncrementIndividualCount ...
func IncrementIndividualCount(svc *dynamodb.DynamoDB, individualType string) error {

	item, err := GetIndividualCount(svc, individualType)

	if err != nil {
		return err
	}
	fmt.Println(item)
	item.Count++
	PutIndividualCount(svc, item)

	fmt.Println("Count "+individualType+": ", item.Count)

	return nil
}

// GetIndividualCount ...
func GetIndividualCount(svc *dynamodb.DynamoDB, individualType string) (IndividualCount, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("individualCount"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(individualType),
			},
		},
	})

	if err != nil {
		return IndividualCount{}, err
	}

	item := IndividualCount{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return IndividualCount{}, err
	}
	return item, nil
}

// PutIndividualCount ...
func PutIndividualCount(svc *dynamodb.DynamoDB, item IndividualCount) {
	itemMarshaled, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		fmt.Println(err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      itemMarshaled,
		TableName: aws.String("individualCount"),
	}
	svc.PutItem(input)
}
