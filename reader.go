package shapefile

import (
	"io"

	"github.com/kellegous/shapefile/dbf"
	"github.com/kellegous/shapefile/shp"
)

// Reader ...
type Reader struct {
	shp *shp.Reader
	dbf *dbf.Reader
}

// Fields ...
func (r *Reader) Fields() []*dbf.Field {
	if r.dbf == nil {
		return nil
	}
	return r.dbf.Fields
}

// FieldCount ...
func (r *Reader) FieldCount() int {
	return len(r.Fields())
}

// ShapeType ...
func (r *Reader) ShapeType() shp.ShapeType {
	return r.shp.ShapeType
}

// BBox ...
func (r *Reader) BBox() *shp.BBox {
	return &r.shp.BBox
}

// Next ...
func (r *Reader) Next() (*Record, error) {
	s, err := r.shp.Next()
	if err != nil {
		return nil, err
	}

	if r.dbf == nil {
		return &Record{Shape: s}, nil
	}

	a, err := r.dbf.Next()
	if err != nil {
		return nil, err
	}

	return &Record{
		Shape:  s,
		Record: a,
	}, nil
}

// NewReader ...
func NewReader(r io.Reader, opts ...Option) (*Reader, error) {
	sr, err := shp.NewReader(r)
	if err != nil {
		return nil, err
	}

	rr := &Reader{
		shp: sr,
	}

	for _, opt := range opts {
		if err := opt(rr); err != nil {
			return nil, err
		}
	}

	return rr, nil
}
