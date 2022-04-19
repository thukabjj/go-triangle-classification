package entrypoint

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
	"github.com/thukabjj/go-triangle-classification/docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func Routes(route *gin.Engine, handlers *entity.Handlers) *gin.RouterGroup {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api"

	apiRoutes := route.Group("/api/triangle")
	apiRoutes.POST("/v1/classifier", handlers.TriangleEntrypoint.TypeIdentifier)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return apiRoutes
}
