package dao

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/rs/xid"
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

type TriangleDao struct {
	ConnectorDynamoDb *infrastructure.ConnectorDynamoDb
}

const tableName = "Triangle"

func NewTriangleDAO(connectorDynamoDb *infrastructure.ConnectorDynamoDb) *TriangleDao {

	result, err := connectorDynamoDb.ListTables()

	shouldCreateTables := true

	if err == nil {

		for _, n := range result.TableNames {
			if *n == tableName {
				shouldCreateTables = false
			}
			fmt.Println("Table ", tableName, "Already Exists: ")
		}

	} else {
		shouldCreateTables = false
		fmt.Println(err.Error())
	}

	if shouldCreateTables {
		createTable(connectorDynamoDb)
	}
	return &TriangleDao{
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
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)
}

func (t *TriangleDao) Store(triangle *domain.Triangle) *domain.Triangle {

	triangleEntity := TriangleEntity{
		ID:           xid.New().String(),
		SideA:        triangle.SideA,
		SideB:        triangle.SideB,
		SideC:        triangle.SideC,
		TriangleType: string(triangle.TriangleType),
	}

	av, err := dynamodbattribute.MarshalMap(triangleEntity)

	if err != nil {
		fmt.Println("Got error calling MarshalMap:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = t.ConnectorDynamoDb.Dynamodb.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling Insert on Table:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	triangle.ID = triangleEntity.ID
	return triangle
}
