package shp

import (
	"math"
	"testing"
)

func polylineMsAreSame(a, b *PolylineM) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfParts == b.NumberOfParts &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allInt32sAreSame(a.Parts, b.Parts) &&
		allPointsAreSame(a.Points, b.Points) &&
		measuresAreSame(&a.M, &b.M)
}

func TestPolylineM(t *testing.T) {
	expectIn(t, "test_files/polylinem", &Expected{
		Header: &Header{
			FileLength: 208,
			ShapeType:  TypePolylineM,
			BBox: BBox{
				MinX: 0,
				MinY: 0,
				MaxX: 89.3,
				MaxY: 90,
			},
		},
		Shapes: []Shape{
			&PolylineM{
				BBox: BBox{
					MinX: 0,
					MinY: 0,
					MaxX: 89.3,
					MaxY: 90,
				},
				NumberOfParts:  2,
				NumberOfPoints: 5,
				Parts:          []int32{0, 3},
				Points: []Point{
					Point{12, 56},
					Point{70, 90},
					Point{0, 0},
					Point{45, 6},
					Point{89.3, 12.2},
				},
				M: M{
					Range: Range{10, 12},
					Measures: []float64{
						10,
						11,
						12,
						math.NaN(),
						math.NaN(),
					},
				},
			},
			&PolylineM{
				BBox: BBox{
					MinX: 0,
					MinY: 0,
					MaxX: 42.42,
					MaxY: 66.66,
				},
				NumberOfParts:  1,
				NumberOfPoints: 2,
				Parts:          []int32{0},
				Points: []Point{
					Point{0, 0},
					Point{42.42, 66.66},
				},
				M: M{
					Range:    Range{4, 5},
					Measures: []float64{4, 5},
				},
			},
		},
	})
}
