package shp

import (
	"math"
	"testing"
)

func pointZsAreSame(a, b *PointZ) bool {
	mEq := (math.IsNaN(a.M) && math.IsNaN(b.M)) ||
		math.Abs(a.M-b.M) < 0.0001
	return math.Abs(a.X-b.X) < 0.0001 &&
		math.Abs(a.Y-b.Y) < 0.0001 &&
		math.Abs(a.Z-b.Z) < 0.0001 &&
		mEq
}
func TestPointZ(t *testing.T) {
	expectIn(t, "../test_files/pointz", &Expected{
		Header: &Header{
			FileLength: 94,
			ShapeType:  TypePointZ,
			BBox:       BBox{0, 0, 10, 20},
		},
		Shapes: []Shape{
			&PointZ{0, 0, 0, 100},
			&PointZ{10, 20, 30, math.NaN()},
		},
	})
}
