package shp

import "testing"

func TestNull(t *testing.T) {
	expectIn(t, "../test_files/null", &Expected{
		Header: &Header{
			FileLength: 62,
			ShapeType:  TypeNull,
			BBox:       BBox{0, 0, 0, 0},
		},
		Shapes: []Shape{
			&Null{},
			&Null{},
		},
	})
}
