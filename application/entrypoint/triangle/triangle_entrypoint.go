package triangle

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle/entity"
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/usercase/triangle/classifier"
)

type TriangleEntrypoint interface {
	TypeIdentifier(ctx *gin.Context)
}

type TriangleEntrypointImpl struct {
	TriangleTypeClassifierUseCase classifier.TriangleTypeClassifierUseCase
}

// Paths Information

// TypeIdentifier godoc
// @Summary Create new videos
// @Description Create a new video
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.TriangleEntrypointRequest true "Create video"
// @Success 200 {object} entity.TriangleEntrypointResponse
// @Failure 400 {object} middleware.Error
// @Failure 401 {object} middleware.Error
// @Router /api/triangle/v1/classifier [post]

func (t *TriangleEntrypointImpl) TypeIdentifier(ctx *gin.Context) {
	request := &entity.TriangleEntrypointRequest{}

	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.Error(err)
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

	ctx.JSON(200, response)
}
