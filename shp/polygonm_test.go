package shp

import (
	"math"
	"testing"
)

func polygonMsAreSame(a, b *PolygonM) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfParts == b.NumberOfParts &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allInt32sAreSame(a.Parts, b.Parts) &&
		allPointsAreSame(a.Points, b.Points) &&
		mdatasAreSame(&a.MData, &b.MData)
}

func TestPolygonM(t *testing.T) {
	expectIn(t, "../test_files/polygonm", &Expected{
		Header: &Header{
			FileLength: 292,
			ShapeType:  TypePolygonM,
			BBox:       BBox{0, 0, 5, 5},
		},
		Shapes: []Shape{
			&PolygonM{
				BBox:           BBox{0, 0, 5, 5},
				NumberOfParts:  2,
				NumberOfPoints: 10,
				Parts:          []int32{0, 5},
				Points: []Point{
					Point{0, 0},
					Point{0, 5},
					Point{5, 5},
					Point{5, 0},
					Point{0, 0},
					Point{1, 1},
					Point{1, 4},
					Point{4, 4},
					Point{4, 1},
					Point{1, 1},
				},
				MData: MData{
					MRange: Range{1, 5},
					M:      []float64{1, 2, 3, 4, 5, math.NaN(), 1, math.NaN(), 2, math.NaN()},
				},
			},
			&PolygonM{
				BBox:           BBox{0, 0, 5, 5},
				NumberOfParts:  1,
				NumberOfPoints: 4,
				Parts:          []int32{0},
				Points: []Point{
					Point{0, 0},
					Point{5, 5},
					Point{5, 0},
					Point{0, 0},
				},
				MData: MData{
					MRange: Range{100, 400.1},
					M:      []float64{100, 200, 300, 400.1},
				},
			},
		},
	})
}

func TestPolygonMOptionalM(t *testing.T) {
	expectIn(t, "../test_files/polygonm_no_m", &Expected{
		Header: &Header{
			FileLength: 220,
			ShapeType:  TypePolygonM,
			BBox:       BBox{0, 0, 5, 5},
		},
		Shapes: []Shape{
			&PolygonM{
				BBox:           BBox{0, 0, 5, 5},
				NumberOfParts:  2,
				NumberOfPoints: 10,
				Parts:          []int32{0, 5},
				Points: []Point{
					Point{0, 0},
					Point{0, 5},
					Point{5, 5},
					Point{5, 0},
					Point{0, 0},
					Point{1, 1},
					Point{1, 4},
					Point{4, 4},
					Point{4, 1},
					Point{1, 1},
				},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M: []float64{
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
					},
				},
			},
			&PolygonM{
				BBox:           BBox{0, 0, 5, 5},
				NumberOfParts:  1,
				NumberOfPoints: 4,
				Parts:          []int32{0},
				Points: []Point{
					Point{0, 0},
					Point{5, 5},
					Point{5, 0},
					Point{0, 0},
				},
				MData: MData{
					MRange: Range{math.NaN(), math.NaN()},
					M: []float64{
						math.NaN(),
						math.NaN(),
						math.NaN(),
						math.NaN(),
					},
				},
			},
		},
	})
}
