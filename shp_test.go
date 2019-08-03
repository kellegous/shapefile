package shp

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"testing"
)

const NoData = -1e39

type Expected struct {
	Header *Header
	Shapes []Shape
}

func float64ToDouble(v float64) float64 {
	if math.IsNaN(v) {
		return NoData
	}
	return v
}

func denormalizeMeasures(m *MData) {
	m.MRange.Min = float64ToDouble(m.MRange.Min)
	m.MRange.Max = float64ToDouble(m.MRange.Max)
	for i, n := 0, len(m.M); i < n; i++ {
		m.M[i] = float64ToDouble(m.M[i])
	}
}

func denormalizeAnyMeasures(data interface{}) {
	shps, ok := data.([]Shape)
	if !ok {
		return
	}

	for _, shp := range shps {
		switch t := shp.(type) {
		case *PointM:
			t.M = float64ToDouble(t.M)
		case *MultiPointM:
			denormalizeMeasures(&t.MData)
		case *PolylineM:
			denormalizeMeasures(&t.MData)
		case *PolygonM:
			denormalizeMeasures(&t.MData)
		case *PointZ:
			t.M = float64ToDouble(t.M)
		case *MultiPointZ:
			denormalizeMeasures(&t.MData)
		case *PolylineZ:
			denormalizeMeasures(&t.MData)
		case *PolygonZ:
			denormalizeMeasures(&t.MData)
		case *MultiPatch:
			denormalizeMeasures(&t.MData)
		}
	}
}

func toJSON(data interface{}) []byte {
	denormalizeAnyMeasures(data)
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return b
}

func bboxesAreSame(a, b *BBox) bool {
	return math.Abs(a.MinX-b.MinX) < 0.0001 &&
		math.Abs(a.MinY-b.MinY) < 0.0001 &&
		math.Abs(a.MaxX-b.MaxX) < 0.0001 &&
		math.Abs(a.MaxY-b.MaxY) < 0.0001
}

func headersAreSame(a, b *Header) bool {
	return a.FileLength == b.FileLength &&
		a.ShapeType == b.ShapeType &&
		bboxesAreSame(&a.BBox, &b.BBox)
}

func allInt32sAreSame(a, b []int32) bool {
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

func allFloat64sAreSame(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := 0, len(a); i < n; i++ {
		if math.Abs(a[i]-b[i]) >= 0.0001 {
			return false
		}
	}

	return true
}

func rangesAreSame(a, b *Range) bool {
	return math.Abs(a.Min-b.Min) < 0.0001 &&
		math.Abs(a.Max-b.Max) < 0.0001
}

func mdataFloatsAreSame(a, b float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	return math.Abs(a-b) < 0.0001
}

func mdatasAreSame(a, b *MData) bool {
	if !mdataFloatsAreSame(a.MRange.Min, b.MRange.Min) ||
		!mdataFloatsAreSame(a.MRange.Max, b.MRange.Max) {
		return false
	}

	if len(a.M) != len(b.M) {
		return false
	}

	for i, n := 0, len(a.M); i < n; i++ {
		if !mdataFloatsAreSame(a.M[i], b.M[i]) {
			return false
		}
	}

	return true
}

func shapesAreSame(a, b Shape) bool {
	switch at := a.(type) {
	case *Null:
		if _, ok := b.(*Null); ok {
			return true
		}
		return false
	case *Point:
		if bt, ok := b.(*Point); ok {
			return pointsAreSame(*at, *bt)
		}
		return false
	case *MultiPoint:
		if bt, ok := b.(*MultiPoint); ok {
			return multiPointAreSame(at, bt)
		}
		return false
	case *Polyline:
		if bt, ok := b.(*Polyline); ok {
			return polylinesAreSame(at, bt)
		}
		return false
	case *Polygon:
		if bt, ok := b.(*Polygon); ok {
			return polygonsAreSame(at, bt)
		}
		return false
	case *PointM:
		if bt, ok := b.(*PointM); ok {
			return pointMsAreSame(at, bt)
		}
		return false
	case *MultiPointM:
		if bt, ok := b.(*MultiPointM); ok {
			return multiPointMsAreSame(at, bt)
		}
		return false
	case *PolylineM:
		if bt, ok := b.(*PolylineM); ok {
			return polylineMsAreSame(at, bt)
		}
		return false
	case *PolygonM:
		if bt, ok := b.(*PolygonM); ok {
			return polygonMsAreSame(at, bt)
		}
		return false
	case *PointZ:
		if bt, ok := b.(*PointZ); ok {
			return pointZsAreSame(at, bt)
		}
		return false
	case *MultiPointZ:
		if bt, ok := b.(*MultiPointZ); ok {
			return multipointZsAreSame(at, bt)
		}
		return false
	case *PolylineZ:
		if bt, ok := b.(*PolylineZ); ok {
			return polylineZsAreSame(at, bt)
		}
		return false
	case *PolygonZ:
		if bt, ok := b.(*PolygonZ); ok {
			return polygonZsAreSame(at, bt)
		}
		return false
	case *MultiPatch:
		if bt, ok := b.(*MultiPatch); ok {
			return multiPatchesAreSame(at, bt)
		}
		return false
	}

	return false
}

func allShapesAreSame(a, b []Shape) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := 0, len(a); i < n; i++ {
		if !shapesAreSame(a[i], b[i]) {
			return false
		}
	}

	return true
}

func expectIn(t *testing.T,
	filename string,
	exp *Expected) {
	r, err := os.Open(filename + ".shp")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	sr, err := NewReader(r)
	if err != nil {
		t.Fatal(err)
	}

	if !headersAreSame(exp.Header, &sr.Header) {
		t.Fatalf("headers expected %s got %s",
			toJSON(exp.Header),
			toJSON(sr.Header))
	}

	var shapes []Shape
	for {
		s, err := sr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		shapes = append(shapes, s)
	}

	if !allShapesAreSame(exp.Shapes, shapes) {
		t.Fatalf("shapes expected %s got %s",
			toJSON(exp.Shapes),
			toJSON(shapes))
	}
}
