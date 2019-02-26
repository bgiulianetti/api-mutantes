package repositories

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/bgiulianetti/api-mutantes/individual"
)

type MockClient struct {
	PutResponse  *dynamodb.PutItemOutput
	GetResponse  *dynamodb.GetItemOutput
	ScanResponse *dynamodb.ScanOutput
	PutError     error
	GetError     error
	ScanError    error
}

func (m MockClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return m.PutResponse, m.PutError
}

func (m MockClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.GetResponse, m.GetError
}

func (m MockClient) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return m.ScanResponse, m.ScanError
}

func TestAddOk(T *testing.T) {

	mutantDNA := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	mutant := individual.Individual{DNA: mutantDNA, ID: "123456"}
	cliente := MockClient{
		GetResponse: &dynamodb.GetItemOutput{Item: map[string]*dynamodb.AttributeValue{}},
		PutResponse: &dynamodb.PutItemOutput{},
		PutError:    nil,
		GetError:    nil,
		ScanError:   nil,
	}
	service, _ := NewPersistenceServiceWithClient(cliente)

	err := service.Add(mutant, "mutante")
	if err != nil {
		T.Error("Error debería ser nil")
	}
}

func TestAddWithPutError(T *testing.T) {

	mutantDNA := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	mutant := individual.Individual{DNA: mutantDNA, ID: "123456"}
	cliente := MockClient{
		PutResponse: &dynamodb.PutItemOutput{},
		PutError:    errors.New("La el tipo de individuo es invalido"),
	}
	service, _ := NewPersistenceServiceWithClient(cliente)

	err := service.Add(mutant, "mutante")
	if err == nil {
		T.Error("El error es nulo")
	}
}

func TestGetStatsWithScanError(T *testing.T) {

	cliente := MockClient{
		GetResponse: &dynamodb.GetItemOutput{Item: map[string]*dynamodb.AttributeValue{}},
		PutResponse: &dynamodb.PutItemOutput{},
		ScanError:   errors.New("falla"),
		PutError:    nil,
	}
	service, _ := NewPersistenceServiceWithClient(cliente)

	_, err := service.GetStats()
	if err == nil {
		T.Error("El error es nulo")
	}
}

func TestGetStatsOk(T *testing.T) {

	statsExpected := individual.Stats{}
	cliente := MockClient{
		ScanResponse: &dynamodb.ScanOutput{Items: []map[string]*dynamodb.AttributeValue{}},
	}
	service, _ := NewPersistenceServiceWithClient(cliente)

	stats, _ := service.GetStats()
	if stats == statsExpected {
		T.Error("El error es nulo")
	}
}

func TestGetStatsOkWithMutant(T *testing.T) {

	mutant := individual.Count{
		ID:    "mutant",
		Count: 0,
	}

	item, _ := dynamodbattribute.MarshalMap(mutant)

	statsExpected := individual.Stats{}
	cliente := MockClient{
		ScanResponse: &dynamodb.ScanOutput{
			Items: []map[string]*dynamodb.AttributeValue{item},
		},
	}
	service, _ := NewPersistenceServiceWithClient(cliente)

	stats, _ := service.GetStats()
	if stats == statsExpected {
		T.Error("El error es nulo")
	}
}

func TestGetStatsOkWithHuman(T *testing.T) {

	human := individual.Count{
		ID:    "human",
		Count: 1,
	}

	item, _ := dynamodbattribute.MarshalMap(human)

	statsExpected := individual.Stats{}
	cliente := MockClient{
		ScanResponse: &dynamodb.ScanOutput{
			Items: []map[string]*dynamodb.AttributeValue{item},
		},
	}
	service, _ := NewPersistenceServiceWithClient(cliente)

	stats, _ := service.GetStats()
	if stats == statsExpected {
		T.Error("El error es nulo")
	}
}
