package shp

import (
	"math"
	"testing"
)

func allPartTypesAreSame(a, b []PartType) bool {
	if len(a) != len(b) {
		return false
	}
	for i, n := 0, len(a); i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func multiPatchesAreSame(a, b *MultiPatch) bool {
	return bboxesAreSame(&a.BBox, &b.BBox) &&
		a.NumberOfParts == b.NumberOfParts &&
		a.NumberOfPoints == b.NumberOfPoints &&
		allPartTypesAreSame(a.PartTypes, b.PartTypes) &&
		allInt32sAreSame(a.Parts, b.Parts) &&
		allPointsAreSame(a.Points, b.Points) &&
		rangesAreSame(&a.ZRange, &b.ZRange) &&
		allFloat64sAreSame(a.Z, b.Z) &&
		mdatasAreSame(&a.MData, &b.MData)
}

func TestMultiPatch(t *testing.T) {
	expectIn(t, "../test_files/multipatch", &Expected{
		Header: &Header{
			FileLength: 386,
			ShapeType:  TypeMultiPatch,
			BBox:       BBox{0, 0, 20, 10},
		},
		Shapes: []Shape{
			&MultiPatch{
				BBox:           BBox{0, 0, 20, 10},
				NumberOfParts:  1,
				NumberOfPoints: 5,
				Parts:          []int32{0},
				PartTypes:      []PartType{TriangleStrip},
				Points: []Point{
					Point{0, 0},
					Point{0, 10},
					Point{10, 10},
					Point{10, 0},
					Point{20, 0},
				},
				ZRange: Range{0, 0},
				Z:      []float64{0, 0, 0, 0, 0},
				MData: MData{
					MRange: Range{12, 16},
					M:      []float64{12, 13, 14, 15, 16},
				},
			},
			&MultiPatch{
				BBox:           BBox{0, 0, 5, 5},
				NumberOfParts:  2,
				NumberOfPoints: 10,
				Parts:          []int32{0, 5},
				PartTypes: []PartType{
					OuterRing,
					Ring,
				},
				Points: []Point{
					Point{0, 0},
					Point{5, 0},
					Point{5, 5},
					Point{0, 5},
					Point{0, 0},
					Point{1, 1},
					Point{4, 1},
					Point{4, 4},
					Point{1, 4},
					Point{1, 1},
				},
				ZRange: Range{1, 1},
				Z:      []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				MData: MData{
					MRange: Range{100, 600},
					M:      []float64{100, 200, 300, 400, 500, 100, 200, 300, 400, 600},
				},
			},
		},
	})
}

func TestMultiPatchOptionalM(t *testing.T) {
	expectIn(t, "../test_files/multipatch_no_m", &Expected{
		Header: &Header{
			FileLength: 310,
			ShapeType:  TypeMultiPatch,
			BBox:       BBox{0, 0, 20, 10},
		},
		Shapes: []Shape{
			&MultiPatch{
				BBox:           BBox{0, 0, 20, 10},
				NumberOfParts:  1,
				NumberOfPoints: 5,
				Parts:          []int32{0},
				PartTypes:      []PartType{TriangleStrip},
				Points: []Point{
					Point{0, 0},
					Point{0, 10},
					Point{10, 10},
					Point{10, 0},
					Point{20, 0},
				},
				ZRange: Range{0, 0},
				Z:      []float64{0, 0, 0, 0, 0},
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
			&MultiPatch{
				BBox:           BBox{0, 0, 5, 5},
				NumberOfParts:  2,
				NumberOfPoints: 10,
				Parts:          []int32{0, 5},
				PartTypes: []PartType{
					OuterRing,
					Ring,
				},
				Points: []Point{
					Point{0, 0},
					Point{5, 0},
					Point{5, 5},
					Point{0, 5},
					Point{0, 0},
					Point{1, 1},
					Point{4, 1},
					Point{4, 4},
					Point{1, 4},
					Point{1, 1},
				},
				ZRange: Range{1, 1},
				Z:      []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
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
		},
	})
}
