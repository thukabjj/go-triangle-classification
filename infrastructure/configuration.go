package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/authentication"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
	authenticationUserCase "github.com/thukabjj/go-triangle-classification/usecase/authentication"

	triangleRepository "github.com/thukabjj/go-triangle-classification/application/database/dynamo/triangle"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle"
	infrastructure "github.com/thukabjj/go-triangle-classification/infrastructure/dynamo"
	"github.com/thukabjj/go-triangle-classification/infrastructure/dynamo/dao"
	"github.com/thukabjj/go-triangle-classification/infrastructure/jwt"
	"github.com/thukabjj/go-triangle-classification/usecase/triangle/classifier"
)

func Run() {
	gin.ForceConsoleColor()

	r := gin.Default()

	r.Use(gin.Recovery(), gin.Logger())

	dynamoDbConnector := infrastructure.NewConnectorDynamoDb()

	entrypoint.Routes(r, buildHandlers(dynamoDbConnector))

	r.Run()
}

func buildHandlers(dynamoDbConnector *infrastructure.ConnectorDynamoDb) *entity.Handlers {
	triangle := InjectTriangleEntrypoint(dynamoDbConnector)
	handlers := &entity.Handlers{
		TriangleEntrypoint: triangle,
		AuthenticationEntrypoint: &authentication.AuthenticationEntrypointImpl{
			AuthenticationUseCase: &authenticationUserCase.AuthenticationUseCaseImpl{
				JwtToken: &jwt.JwtTokenImpl{},
			},
		},
	}
	return handlers
}

func InjectTriangleEntrypoint(dynamoDbConnector *infrastructure.ConnectorDynamoDb) triangle.TriangleEntrypoint {

	triangle := &triangle.TriangleEntrypointImpl{
		TriangleTypeClassifierUseCase: &classifier.TriangleTypeClassifierUseCaseImpl{
			TriangleRepository: &triangleRepository.TriangleRepositoryImpl{
				TriangleDao: dao.NewTriangleDAOImpl(dynamoDbConnector),
			},
		},
	}
	return triangle
}
