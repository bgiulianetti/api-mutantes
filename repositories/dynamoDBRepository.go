package repositories

import (
	"math"

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
	sess, _ := session.NewSession(&aws.Config{Region: aws.String("sa-east-1")})
	return PersistenceService{Session: dynamodb.New(sess)}, nil
}

// Add agrega un mutante a dynamodb
func (p PersistenceService) Add(individualToPersist individual.Individual, individualType string) error {

	individualToPersist.ID = utils.ConcatenateStringArray(individualToPersist.DNA)

	item, _ := dynamodbattribute.MarshalMap(individualToPersist)

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(individualType),
	}

	_, err := p.Session.PutItem(input)
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
	_ = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)

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
		ratio := response.CountMutant / response.CountHuman
		response.Ratio = math.Round(ratio*100) / 100
	}
	return response, nil
}
