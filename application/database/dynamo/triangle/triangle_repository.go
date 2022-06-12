package triangle

import (
	"github.com/thukabjj/go-triangle-classification/domain"
)

type TriangleRepositoryImpl struct {
	TriangleDao TriangleDao
}

func (t *TriangleRepositoryImpl) Store(triangle *domain.Triangle) *domain.Triangle {
	triangleSaved := t.TriangleDao.Save(triangle)
	return triangleSaved
}
