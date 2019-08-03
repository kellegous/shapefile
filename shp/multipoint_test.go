package shp

import "testing"

func multiPointAreSame(a, b *MultiPoint) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allPointsAreSame(a.Points, b.Points)
}

func TestMultiPoint(t *testing.T) {
	expectIn(t, "../test_files/multipoint", &Expected{
		Header: &Header{
			FileLength: 146,
			ShapeType:  TypeMultiPoint,
			BBox:       BBox{0, 0, 20, 10},
		},
		Shapes: []Shape{
			&MultiPoint{
				BBox: BBox{
					MinX: 0,
					MinY: 0,
					MaxX: 20,
					MaxY: 10,
				},
				NumberOfPoints: 4,
				Points: []Point{
					Point{0, 0},
					Point{10, 10},
					Point{20, 10},
					Point{5, 5},
				},
			},
			&MultiPoint{
				BBox: BBox{
					MinX: 0,
					MinY: 0,
					MaxX: 4,
					MaxY: 5,
				},
				NumberOfPoints: 2,
				Points: []Point{
					Point{0, 0},
					Point{4, 5},
				},
			},
		},
	})
}
