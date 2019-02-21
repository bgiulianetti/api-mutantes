package repositories

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bgiulianetti/api-mutantes/individual"
	"github.com/bgiulianetti/api-mutantes/utils"
)

// PersistenceService ...
type PersistenceService struct {
	Session Client
}

// Client ...
type Client interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

// NewPersistenceServiceWithClient crea una sesion de conexión a dynamodb
func NewPersistenceServiceWithClient(cliente Client) (PersistenceService, error) {
	return PersistenceService{Session: cliente}, nil
}

// NewPersistenceService crea una sesion de conexión a dynamodb
func NewPersistenceService() (PersistenceService, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("sa-east-1")})
	if err != nil {
		return PersistenceService{}, err
	}
	svc := dynamodb.New(sess)

	return PersistenceService{Session: svc}, nil
}

// Add agrega un mutante a dynamodb
func (p PersistenceService) Add(individualToPersist individual.Individual, individualType string) error {

	individualToPersist.ID = utils.ConcatenateStringArray(individualToPersist.DNA)

	item, err := dynamodbattribute.MarshalMap(individualToPersist)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(individualType),
	}

	_, err = p.Session.PutItem(input)
	if err != nil {
		return err
	}

	err = p.IncrementCount(individualType)
	if err != nil {
		return err
	}
	return nil
}

// GetStats ...
func (p PersistenceService) GetStats() (individual.Stats, error) {

	result, err := p.Session.Scan(&dynamodb.ScanInput{
		TableName: aws.String("individualCount"),
	})
	if err != nil {
		return individual.Stats{}, err
	}

	items := []individual.Count{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return individual.Stats{}, err
	}
	response := individual.Stats{}

	for _, item := range items {
		if item.ID == "human" {
			response.CountHuman = float64(item.Count)
		}
		if item.ID == "mutant" {
			response.CountMutant = float64(item.Count)
		}
	}
	if response.CountHuman == 0 {
		response.Ratio = 1
	} else {
		response.Ratio = response.CountMutant / response.CountHuman
	}
	return response, nil
}

// IncrementCount ...
func (p PersistenceService) IncrementCount(individualType string) error {

	item, err := p.GetCount(individualType)
	if err != nil {
		return err
	}
	fmt.Println(item)
	item.Count++
	err = p.PutIndividualCount(item)
	if err != nil {
		return err
	}
	fmt.Println("Count "+individualType+": ", item.Count)
	return nil
}

// GetCount ...
func (p PersistenceService) GetCount(individualType string) (individual.Count, error) {

	tabla := "individualCount"
	result, err := p.Session.GetItem(&dynamodb.GetItemInput{
		TableName: &tabla,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(individualType),
			},
		},
	})

	if err != nil {
		return individual.Count{}, err
	}

	count := individual.Count{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &count)
	if err != nil {
		return individual.Count{}, err
	}
	return count, nil
}

// PutIndividualCount ...
func (p PersistenceService) PutIndividualCount(count individual.Count) error {

	itemMarshaled, err := dynamodbattribute.MarshalMap(count)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      itemMarshaled,
		TableName: aws.String("individualCount"),
	}
	p.Session.PutItem(input)
	return nil
}
