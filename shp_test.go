package shp

import (
	"encoding/json"
	"io"
	"math"
	"os"
	"testing"
)

type Expected struct {
	Header *Header
	Shapes []Shape
}

func toJSON(data interface{}) []byte {
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
