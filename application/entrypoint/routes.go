package entrypoint

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
)

func Routes(route *gin.Engine, handlers *entity.Handlers) *gin.RouterGroup {

	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootRoute := route.Group("")
	{
		triangleRoutes := route.Group("/api/triangle")
		{
			triangleRoutes.POST("/v1/classifier", handlers.TriangleEntrypoint.TypeIdentifier)
		}

		authRoutes := route.Group("/auth")
		{
			authRoutes.POST("/login", handlers.AuthenticationEntrypoint.Login)
		}

	}
	return rootRoute
}
