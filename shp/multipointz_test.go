package shp

import (
	"math"
	"testing"
)

func multipointZsAreSame(a, b *MultiPointZ) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allPointsAreSame(a.Points, b.Points) &&
		rangesAreSame(&a.ZRange, &b.ZRange) &&
		allFloat64sAreSame(a.Z, b.Z) &&
		mdatasAreSame(&a.MData, &b.MData)
}

func TestMultiPointZ(t *testing.T) {
	expectIn(t, "../test_files/multipointz", &Expected{
		Header: &Header{
			FileLength: 226,
			ShapeType:  TypeMultiPointZ,
			BBox:       BBox{0, 0, 20, 10},
		},
		Shapes: []Shape{
			&MultiPointZ{
				BBox:           BBox{0, 0, 20, 10},
				NumberOfPoints: 4,
				Points: []Point{
					Point{0, 0},
					Point{10, 10},
					Point{20, 10},
					Point{5, 5},
				},
				ZRange: Range{0, 3},
				Z:      []float64{0, 1, 2, 3},
				MData: MData{
					MRange: Range{100, 400},
					M:      []float64{100, 200, math.NaN(), 400},
				},
			},
			&MultiPointZ{
				BBox:           BBox{0, 0, 4, 5},
				NumberOfPoints: 2,
				Points: []Point{
					Point{0, 0},
					Point{4, 5},
				},
				ZRange: Range{0, 6},
				Z:      []float64{0, 6},
				MData: MData{
					MRange: Range{12.2, 13.3},
					M:      []float64{12.2, 13.3},
				},
			},
		},
	})
}

func TestMultiPointZOptionalM(t *testing.T) {
	expectIn(t, "../test_files/multipointz_no_m", &Expected{
		Header: &Header{
			FileLength: 186,
			ShapeType:  TypeMultiPointZ,
			BBox:       BBox{0, 0, 20, 10},
		},
		Shapes: []Shape{
			&MultiPointZ{
				BBox:           BBox{0, 0, 20, 10},
				NumberOfPoints: 4,
				Points: []Point{
					Point{0, 0},
					Point{10, 10},
					Point{20, 10},
					Point{5, 5},
				},
				ZRange: Range{0, 3},
				Z:      []float64{0, 1, 2, 3},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M:      []float64{math.NaN(), math.NaN(), math.NaN(), math.NaN()},
				},
			},
			&MultiPointZ{
				BBox:           BBox{0, 0, 4, 5},
				NumberOfPoints: 2,
				Points: []Point{
					Point{0, 0},
					Point{4, 5},
				},
				ZRange: Range{0, 6},
				Z:      []float64{0, 6},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M:      []float64{math.NaN(), math.NaN()},
				},
			},
		},
	})
}
