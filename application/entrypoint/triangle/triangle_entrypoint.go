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

// Upload example
// @Summary Upload file
// @Description Upload file
// @ID file.upload
// @Accept  multipart/form-data
// @Produce  json
// @Param   file formData file true  "this is a test file"
// @Success 200 {string} string "ok"
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
