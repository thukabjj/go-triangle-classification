package entrypoint

import (
	"github.com/gin-gonic/gin"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
)

func Routes(route *gin.Engine, handlers *entity.Handlers) *gin.RouterGroup {

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
