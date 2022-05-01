package repository

import "github.com/thukabjj/go-triangle-classification/domain"

type TriangleRepository interface {
	Store(triangle *domain.Triangle) *domain.Triangle
}
