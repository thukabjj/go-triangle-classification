package triangle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle/entity"
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/usecase/triangle/classifier"
)

type TriangleEntrypoint interface {
	TypeIdentifier(ctx *gin.Context)
}

type TriangleEntrypointImpl struct {
	TriangleTypeClassifierUseCase classifier.TriangleTypeClassifierUseCase
}

// PostTriangleIdentify             godoc
// @Summary      Identify a tiangle type
// @Description  Takes the triangle request JSON and identifies the triangle's type, and stores it in DB. Return saved JSON.
// @Tags         triangle
// @Produce      json
// @Param        TraingleRequest body      entity.TriangleEntrypointRequest  true  "TriangleEntrypointRequest JSON information"
// @Security     ApiKeyAuth
// @Success      200   {object}  entity.TriangleEntrypointResponse
// @Failure      401   {object}  middleware.Error "User not authorized!"
// @Failure      422   {object}  middleware.Error "Unprocessable Entity"
// @Router       /api/triangle/v1/classifier [post]
func (t *TriangleEntrypointImpl) TypeIdentifier(ctx *gin.Context) {
	request := &entity.TriangleEntrypointRequest{}

	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.Error(err)
		ctx.Abort()
		return
	}
	triangle := &domain.Triangle{
		SideA: *request.SideA,
		SideB: *request.SideB,
		SideC: *request.SideC,
	}
	triangleType := t.TriangleTypeClassifierUseCase.Execute(triangle)
	response := &entity.TriangleEntrypointResponse{
		TriangleType: string(triangleType),
	}

	ctx.JSON(http.StatusOK, response)
}
