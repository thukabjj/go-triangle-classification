package entity

type TriangleEntity struct {
	ID           string  `json:"id"`
	SideA        float64 `json:"sideA"`
	SideB        float64 `json:"sideB"`
	SideC        float64 `json:"sideC"`
	TriangleType string  `json:"triangleType"`
}
