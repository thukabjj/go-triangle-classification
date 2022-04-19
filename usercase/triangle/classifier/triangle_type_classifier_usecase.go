package classifier

import (
	"fmt"

	"github.com/thukabjj/go-triangle-classification/domain"
	repository "github.com/thukabjj/go-triangle-classification/usercase/gateway/repository/triangle"
)

type TriangleTypeClassifierUseCase interface {
	Execute(triangle *domain.Triangle) domain.TriangleType
}

type TriangleTypeClassifierUseCaseImpl struct {
	TriangleRepository repository.TriangleRepository
}

func (t *TriangleTypeClassifierUseCaseImpl) Execute(triangle *domain.Triangle) domain.TriangleType {
	triangleType := triangle.IdentifyTriangleType()
	triangle.TriangleType = triangleType
	savedTriangle := t.TriangleRepository.Store(triangle)
	fmt.Println(savedTriangle)
	return savedTriangle.TriangleType
}
