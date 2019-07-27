package shp

import (
	"math"
	"testing"
)

func pointsAreSame(a, b *Point) bool {
	return math.Abs(a.X-b.X) < 0.0001 &&
		math.Abs(a.Y-b.Y) < 0.0001
}

func TestPoint(t *testing.T) {
	expectIn(t, "test_files/point", &Expected{
		Header: &Header{
			FileLength: 78,
			ShapeType:  TypePoint,
			BBox:       BBox{10, 20, 50, 40},
		},
		Shapes: []Shape{
			&Point{X: 10, Y: 20.0},
			&Point{X: 50, Y: 40},
		},
	})
}
