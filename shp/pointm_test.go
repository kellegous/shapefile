package shp

import (
	"math"
	"testing"
)

func pointMsAreSame(a, b *PointM) bool {
	mEq := (math.IsNaN(a.M) && math.IsNaN(b.M)) ||
		math.Abs(a.M-b.M) < 0.0001
	return math.Abs(a.X-b.X) < 0.0001 &&
		math.Abs(a.Y-b.Y) < 0.0001 &&
		mEq
}

func TestPointM(t *testing.T) {
	expectIn(t, "../test_files/pointm", &Expected{
		Header: &Header{
			FileLength: 104,
			ShapeType:  TypePointM,
			BBox: BBox{
				MinX: 0,
				MinY: 0,
				MaxX: 42.2,
				MaxY: 55.5,
			},
		},
		Shapes: []Shape{
			&PointM{0, 0, 100},
			&PointM{42.2, 55.5, 50},
			&PointM{20.3, 30, math.NaN()},
		},
	})
}
