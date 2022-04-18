package entrypoint

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/entity"
	"github.com/thukabjj/go-triangle-classification/docs"
	// swagger embed files
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/triangle

// @securityDefinitions.basic  BasicAuth

func Routes(route *gin.Engine, handlers *entity.Handlers) *gin.RouterGroup {
	docs.SwaggerInfo.BasePath = "/api/triangle"
	apiRoutes := route.Group("/api/triangle")
	apiRoutes.POST("/v1/classifier", handlers.TriangleEntrypoint.TypeIdentifier)
	apiRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return apiRoutes
}
