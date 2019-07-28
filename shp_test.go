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

func denormalizeMeasures(m *M) {
	m.Range.Min = float64ToDouble(m.Range.Min)
	m.Range.Max = float64ToDouble(m.Range.Max)
	for i, n := 0, len(m.Measures); i < n; i++ {
		m.Measures[i] = float64ToDouble(m.Measures[i])
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
			denormalizeMeasures(&t.M)
		case *PolylineM:
			denormalizeMeasures(&t.M)
		case *PolygonM:
			denormalizeMeasures(&t.M)
		case *PointZ:
			t.M = float64ToDouble(t.M)
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

func float64sAreSame(a, b float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	return math.Abs(a-b) < 0.0001
}

func measuresAreSame(a, b *M) bool {
	if !float64sAreSame(a.Range.Min, b.Range.Min) ||
		!float64sAreSame(a.Range.Max, b.Range.Max) {
		return false
	}

	if len(a.Measures) != len(b.Measures) {
		return false
	}

	for i, n := 0, len(a.Measures); i < n; i++ {
		if !float64sAreSame(a.Measures[i], b.Measures[i]) {
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
