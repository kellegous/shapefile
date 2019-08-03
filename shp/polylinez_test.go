package shp

import (
	"math"
	"testing"
)

func polylineZsAreSame(a, b *PolylineZ) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfParts == b.NumberOfParts &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allInt32sAreSame(a.Parts, b.Parts) &&
		allPointsAreSame(a.Points, b.Points) &&
		rangesAreSame(&a.ZRange, &b.ZRange) &&
		allFloat64sAreSame(a.Z, b.Z) &&
		mdatasAreSame(&a.MData, &b.MData)
}

func TestPolylineZ(t *testing.T) {
	expectIn(t, "../test_files/polylinez", &Expected{
		Header: &Header{
			FileLength: 252,
			ShapeType:  TypePolylineZ,
			BBox:       BBox{0, 0, 89.3, 90},
		},
		Shapes: []Shape{
			&PolylineZ{
				BBox:           BBox{0, 0, 89.3, 90},
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
				ZRange: Range{11.1, 55.5},
				Z:      []float64{11.1, 22.2, 33.3, 44.4, 55.5},
				MData: MData{
					MRange: Range{10, 12},
					M:      []float64{10, 11, 12, math.NaN(), math.NaN()},
				},
			},
			&PolylineZ{
				BBox:           BBox{0, 0, 42.42, 66.66},
				NumberOfParts:  1,
				NumberOfPoints: 2,
				Parts:          []int32{0},
				Points: []Point{
					Point{0, 0},
					Point{42.42, 66.66},
				},
				ZRange: Range{66.6, 77.7},
				Z:      []float64{66.6, 77.7},
				MData: MData{
					MRange: Range{4, 5},
					M:      []float64{4, 5},
				},
			},
		},
	})
}

func TestPolylineZOptionalM(t *testing.T) {
	expectIn(t, "../test_files/polylinez_no_m", &Expected{
		Header: &Header{
			FileLength: 208,
			ShapeType:  TypePolylineZ,
			BBox:       BBox{0, 0, 89.3, 90},
		},
		Shapes: []Shape{
			&PolylineZ{
				BBox:           BBox{0, 0, 89.3, 90},
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
				ZRange: Range{11.1, 55.5},
				Z:      []float64{11.1, 22.2, 33.3, 44.4, 55.5},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M: []float64{
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
					},
				},
			},
			&PolylineZ{
				BBox:           BBox{0, 0, 42.42, 66.66},
				NumberOfParts:  1,
				NumberOfPoints: 2,
				Parts:          []int32{0},
				Points: []Point{
					Point{0, 0},
					Point{42.42, 66.66},
				},
				ZRange: Range{66.6, 77.7},
				Z:      []float64{66.6, 77.7},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M: []float64{
						math.NaN(),
						math.NaN(),
					},
				},
			},
		},
	})
}
