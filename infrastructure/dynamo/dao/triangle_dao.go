package dao

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/thukabjj/go-triangle-classification/domain"
	infrastructure "github.com/thukabjj/go-triangle-classification/infrastructure/dynamo"
)

type TriangleEntity struct {
	ID           string  `json:"id"`
	SideA        float64 `json:"sideA"`
	SideB        float64 `json:"sideB"`
	SideC        float64 `json:"sideC"`
	TriangleType string  `json:"triangleType"`
}

type TriangleDaoImpl struct {
	ConnectorDynamoDb *infrastructure.ConnectorDynamoDb
}

const tableName = "Triangle"

func NewTriangleDAOImpl(connectorDynamoDb *infrastructure.ConnectorDynamoDb) *TriangleDaoImpl {

	result, err := connectorDynamoDb.ListTables()

	shouldCreateTables := true

	if err == nil {

		for _, n := range result.TableNames {
			if *n == tableName {
				shouldCreateTables = false
			}
			log.Printf("Table %s already exists!", tableName)
		}

	} else {
		shouldCreateTables = false
		log.Fatalf("Got error calling ListTables: %s", err.Error())
	}

	if shouldCreateTables {
		createTable(connectorDynamoDb)
	}
	return &TriangleDaoImpl{
		ConnectorDynamoDb: connectorDynamoDb,
	}
}

func createTable(connectorDynamoDb *infrastructure.ConnectorDynamoDb) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String("PAY_PER_REQUEST"),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := connectorDynamoDb.Dynamodb.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err.Error())
	}

	log.Printf("Table %s was created!", tableName)
}

func (t *TriangleDaoImpl) Save(triangle *domain.Triangle) *domain.Triangle {

	triangleEntity := TriangleEntity{
		ID:           uuid.New().String(),
		SideA:        triangle.SideA,
		SideB:        triangle.SideB,
		SideC:        triangle.SideC,
		TriangleType: string(triangle.TriangleType),
	}

	av, err := dynamodbattribute.MarshalMap(triangleEntity)

	if err != nil {
		log.Fatalf("Got error calling MarshalMap: %s", err.Error())
		return nil
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = t.ConnectorDynamoDb.Dynamodb.PutItem(input)

	if err != nil {
		log.Fatalf("Got error calling Insert on Table: %s", err.Error())
		return nil
	}
	triangle.ID = triangleEntity.ID
	return triangle
}
