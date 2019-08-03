package shp

import (
	"math"
	"testing"
)

func multiPointMsAreSame(a, b *MultiPointM) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allPointsAreSame(a.Points, b.Points) &&
		mdatasAreSame(&a.MData, &b.MData)
}

func TestMultiPointM(t *testing.T) {
	expectIn(t, "../test_files/multipointm", &Expected{
		Header: &Header{
			FileLength: 162,
			ShapeType:  TypeMultiPointM,
			BBox:       BBox{0, 0, 3, 5},
		},
		Shapes: []Shape{
			&MultiPointM{
				BBox:           BBox{0, 0, 3, 5},
				NumberOfPoints: 3,
				Points: []Point{
					Point{0, 0},
					Point{1, 2},
					Point{3, 5},
				},
				MData: MData{
					MRange: Range{4, 8},
					M:      []float64{4, 8, math.NaN()},
				},
			},
			&MultiPointM{
				BBox:           BBox{0, 1, 0, 1},
				NumberOfPoints: 1,
				Points: []Point{
					Point{0, 1},
				},
				MData: MData{
					MRange: Range{100, 100},
					M:      []float64{100},
				},
			},
		},
	})
}

func TestMultiPointMOptionalM(t *testing.T) {
	expectIn(t, "../test_files/multipointm_no_m", &Expected{
		Header: &Header{
			FileLength: 130,
			ShapeType:  TypeMultiPointM,
			BBox:       BBox{0, 0, 3, 5},
		},
		Shapes: []Shape{
			&MultiPointM{
				BBox:           BBox{0, 0, 3, 5},
				NumberOfPoints: 3,
				Points: []Point{
					Point{0, 0},
					Point{1, 2},
					Point{3, 5},
				},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M:      []float64{math.NaN(), math.NaN(), math.NaN()},
				},
			},
			&MultiPointM{
				BBox:           BBox{0, 1, 0, 1},
				NumberOfPoints: 1,
				Points: []Point{
					Point{0, 1},
				},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M:      []float64{math.NaN()},
				},
			},
		},
	})
}
