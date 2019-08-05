package shapefile

import (
	"io"
	"os"
	"testing"

	"github.com/kellegous/shapefile/shp"
)

func TestNextWithDBF(t *testing.T) {
	sr, err := os.Open("test_files/multipatch.shp")
	if err != nil {
		t.Fatal(err)
	}
	defer sr.Close()

	dr, err := os.Open("test_files/multipatch.dbf")
	if err != nil {
		t.Fatal(err)
	}
	defer dr.Close()

	r, err := NewReader(sr, WithDBF(dr))
	if err != nil {
		t.Fatal(err)
	}

	if r.FieldCount() != 6 {
		t.Fatalf("expected 6 fields, got %d", r.FieldCount())
	}

	if r.ShapeType() != shp.TypeMultiPatch {
		t.Fatalf("expected ShapeType of MultiPatch got %s", r.ShapeType())
	}

	bbox := shp.BBox{MinX: 0, MinY: 0, MaxX: 20, MaxY: 10}
	if *r.BBox() != bbox {
		t.Fatalf("expected BBox of %#v got %#v", bbox, r.BBox())
	}

	for i := 0; i < 2; i++ {
		rec, err := r.Next()
		if err != nil {
			t.Fatal(err)
		}

		if _, ok := rec.Shape.(*shp.MultiPatch); !ok {
			t.Fatalf("expected a MultiPatch got %#v", rec.Shape)
		}

		if len(rec.Attrs()) != 6 {
			t.Fatalf("expected 6 attrs got %d", len(rec.Attrs()))
		}
	}

	if _, err = r.Next(); err != io.EOF {
		t.Fatalf("expected EOF but got %s", err)
	}
}

func TestNextWithoutDBF(t *testing.T) {
	sr, err := os.Open("test_files/multipatch.shp")
	if err != nil {
		t.Fatal(err)
	}
	defer sr.Close()

	r, err := NewReader(sr)
	if err != nil {
		t.Fatal(err)
	}

	if r.FieldCount() != 0 {
		t.Fatalf("expected 0 fields, got %d", r.FieldCount())
	}

	if r.ShapeType() != shp.TypeMultiPatch {
		t.Fatalf("expected ShapeType of MultiPatch got %s", r.ShapeType())
	}

	bbox := shp.BBox{MinX: 0, MinY: 0, MaxX: 20, MaxY: 10}
	if *r.BBox() != bbox {
		t.Fatalf("expected BBox of %#v got %#v", bbox, r.BBox())
	}

	for i := 0; i < 2; i++ {
		rec, err := r.Next()
		if err != nil {
			t.Fatal(err)
		}

		if _, ok := rec.Shape.(*shp.MultiPatch); !ok {
			t.Fatalf("expected a MultiPatch got %#v", rec.Shape)
		}

		if len(rec.Attrs()) != 0 {
			t.Fatalf("expected 0 attrs got %d", len(rec.Attrs()))
		}
	}

	if _, err = r.Next(); err != io.EOF {
		t.Fatalf("expected EOF but got %s", err)
	}
}

func TestNextUnexpectedEOF(t *testing.T) {
	// has 3 shapes
	sr, err := os.Open("test_files/pointm.shp")
	if err != nil {
		t.Fatal(err)
	}
	defer sr.Close()

	// only has 2 records
	dr, err := os.Open("test_files/multipatch.dbf")
	if err != nil {
		t.Fatal(err)
	}
	defer dr.Close()

	r, err := NewReader(sr, WithDBF(dr))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 2; i++ {
		if _, err := r.Next(); err != nil {
			t.Fatal(err)
		}
	}

	if _, err := r.Next(); err != io.ErrUnexpectedEOF {
		t.Fatalf("expected ErrUnexpectedEOF got %s", err)
	}
}
