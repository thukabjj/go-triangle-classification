package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle"
	"github.com/thukabjj/go-triangle-classification/application/middleware"
	infrastructure "github.com/thukabjj/go-triangle-classification/infrastructure/dynamo"
	"github.com/thukabjj/go-triangle-classification/infrastructure/dynamo/dao"
	"github.com/thukabjj/go-triangle-classification/usercase/triangle/classifier"
)

func Run() {
	gin.ForceConsoleColor()

	r := gin.Default()

	r.Use(gin.Recovery(), gin.Logger(), middleware.ErrorHandler)

	dynamoDbConnector := infrastructure.NewConnectorDynamoDb()

	entrypoint.Routes(r, buildHandlers(dynamoDbConnector))

	r.Run()
}

func buildHandlers(dynamoDbConnector *infrastructure.ConnectorDynamoDb) *entity.Handlers {
	triangle := InjectTriangleEntrypoint(dynamoDbConnector)
	handlers := &entity.Handlers{
		TriangleEntrypoint: triangle,
	}
	return handlers
}

func InjectTriangleEntrypoint(dynamoDbConnector *infrastructure.ConnectorDynamoDb) triangle.TriangleEntrypoint {

	triangle := &triangle.TriangleEntrypointImpl{
		TriangleTypeClassifierUseCase: &classifier.TriangleTypeClassifierUseCaseImpl{
			TriangleRepository: dao.NewTriangleDAO(dynamoDbConnector),
		},
	}
	return triangle
}
