package triangle

import "github.com/thukabjj/go-triangle-classification/domain"

type TriangleDao interface {
	Save(triangle *domain.Triangle) *domain.Triangle
}
