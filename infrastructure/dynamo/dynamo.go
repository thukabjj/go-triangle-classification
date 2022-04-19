package infrastructure

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ConnectorDynamoDb struct {
	Dynamodb *dynamodb.DynamoDB
}

func NewConnectorDynamoDb() *ConnectorDynamoDb {

	region := "us-east-1"
	profile := "localstack"

	session, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:                        aws.String(region),
			Endpoint:                      aws.String("http://127.0.0.1:4566"),
			Credentials:                   credentials.NewStaticCredentials("fakeKeyLocalStack", "fakeSecretLocalStack", ""),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
		Profile: profile,
	})

	if err != nil {
		log.Fatalf("Got error on creating DynamoDB session connector: %s", err.Error())
	}

	svc := dynamodb.New(session)

	return &ConnectorDynamoDb{
		Dynamodb: svc,
	}
}

func (c *ConnectorDynamoDb) ListTables() (*dynamodb.ListTablesOutput, error) {
	tables := &dynamodb.ListTablesInput{}
	result, err := c.Dynamodb.ListTables(tables)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				log.Fatalf("Got error on listing tables DynamoDB: %s -  %s", dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				log.Fatalf("Unexpected Dynamodb error: %s", aerr.Error())
			}
			return nil, err
		}

	}
	return result, nil
}
