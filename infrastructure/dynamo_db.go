package infrastructure

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/thukabjj/go-triangle-classification/domain"
)

type DynamoDbRepository struct {
}

func NewDynamoDbRepo() *DynamoDbRepository {

	region := "us-east-1"
	profile := "localstack"
	//Create Table
	tableName := "Triangle"

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(region),
			Endpoint:                      aws.String("http://127.0.0.1:4566"),
			Credentials:                   credentials.NewStaticCredentials("fakeKeyLocalStack", "fakeSecretLocalStack", ""),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
		Profile: profile,
	})

	if err != nil {
		fmt.Println(err)
	}

	svc := dynamodb.New(sess)
	tables := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(tables)
	// Get the list of tables

	createTable := true
	if err != nil {
		createTable = false
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())

		}

		for _, n := range result.TableNames {
			if *n == tableName {
				createTable = false
			}
			fmt.Println(*n)
		}

		if createTable {
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

			_, err = svc.CreateTable(input)
			if err != nil {
				log.Fatalf("Got error calling CreateTable: %s", err)
			}

			fmt.Println("Created the table", tableName)
		}

	}

	return nil
}
func (r *DynamoDbRepository) Store(triangle *domain.Triangle) *domain.Triangle {

	return triangle
}
