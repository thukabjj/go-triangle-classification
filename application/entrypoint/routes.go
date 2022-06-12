package entrypoint

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
	"github.com/thukabjj/go-triangle-classification/application/middleware"
)

func Routes(route *gin.Engine, handlers *entity.Handlers) *gin.RouterGroup {

	route.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.Use(middleware.ValidationErrorsMiddleware)
	rootRoute := route.Group("")
	{
		triangleRoutes := route.Group("/api/triangle")
		triangleRoutes.Use(middleware.ValidateToken)
		{
			triangleRoutes.POST("/v1/classifier", handlers.TriangleEntrypoint.TypeIdentifier)
		}

		authRoutes := route.Group("/auth")
		authRoutes.Use(middleware.JwtValidationsError)
		{
			authRoutes.POST("/login", handlers.AuthenticationEntrypoint.Login)
		}

	}
	return rootRoute
}
