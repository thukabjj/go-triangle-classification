package repository

import "github.com/thukabjj/go-triangle-classification/domain"

// TriangleRepository handles the persistence of triangles.
type TriangleRepository interface {
	Store(triangle *domain.Triangle) *domain.Triangle
}
