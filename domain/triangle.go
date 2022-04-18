package domain

type TriangleType string

const (
	TriangleTypeEquilateral TriangleType = "equilateral"
	TriangleTypeIsosceles   TriangleType = "isosceles"
	TriangleTypeScalene     TriangleType = "scalene"
	TriangleTypeUndefined   TriangleType = "undefined"
)

func (t TriangleType) GetTriangleType(typeTriangle string) string {
	switch typeTriangle {

	case string(TriangleTypeEquilateral):
		return string(TriangleTypeEquilateral)
	case string(TriangleTypeIsosceles):
		return string(TriangleTypeIsosceles)
	case string(TriangleTypeScalene):
		return string(TriangleTypeScalene)
	default:
		return string(TriangleTypeUndefined)

	}
}

type Triangle struct {
	ID           string
	SideA        float64
	SideB        float64
	SideC        float64
	TriangleType TriangleType
}

func (t *Triangle) IdentifyTriangleType() TriangleType {
	if t.SideA == t.SideB && t.SideA == t.SideC {
		return TriangleTypeEquilateral
	}
	if t.SideA == t.SideB || t.SideB == t.SideC || t.SideC == t.SideA {
		return TriangleTypeIsosceles
	}
	return TriangleTypeScalene
}
