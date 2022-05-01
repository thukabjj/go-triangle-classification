package entity

import (
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/authentication"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/triangle"
)

type Handlers struct {
	TriangleEntrypoint       triangle.TriangleEntrypoint
	AuthenticationEntrypoint authentication.AuthenticationEntrypoint
}
