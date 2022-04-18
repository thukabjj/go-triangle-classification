package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle"
	"github.com/thukabjj/go-triangle-classification/application/middleware"
	"github.com/thukabjj/go-triangle-classification/usercase/triangle/classifier"
)

func Run() {
	gin.ForceConsoleColor()

	r := gin.Default()

	r.Use(gin.Recovery(), gin.Logger(), middleware.ErrorHandler)

	entrypoint.Routes(r, buildHandlers())

	r.Run()
}

func buildHandlers() *entity.Handlers {
	triangle := InjectTriangleEntrypoint()
	handlers := &entity.Handlers{
		TriangleEntrypoint: triangle,
	}
	return handlers
}

func InjectTriangleEntrypoint() triangle.TriangleEntrypoint {

	triangle := &triangle.TriangleEntrypointImpl{
		TriangleTypeClassifierUseCase: &classifier.TriangleTypeClassifierUseCaseImpl{
			TriangleRepository: NewDynamoDbRepo(),
		},
	}
	return triangle
}
