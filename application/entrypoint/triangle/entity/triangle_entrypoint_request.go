package entity

type TriangleEntrypointRequest struct {
	SideA *float64 `json:"sideA" binding:"required,omitempty,gt=0"`
	SideB *float64 `json:"sideB" binding:"required,omitempty,gt=0"`
	SideC *float64 `json:"sideC" binding:"required,omitempty,gt=0"`
}
