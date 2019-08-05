package shapefile

import (
	"io"

	"github.com/kellegous/shapefile/dbf"
	"github.com/kellegous/shapefile/shp"
)

// Reader provides sequential access to the records in the shapefile.
type Reader struct {
	shp *shp.Reader
	dbf *dbf.Reader
}

// Fields returns the fields that were declared in the dbf header structure. If no
// dbf was provided when the Reader was created, this will return nil.
func (r *Reader) Fields() []*dbf.Field {
	if r.dbf == nil {
		return nil
	}
	return r.dbf.Fields
}

// FieldCount returns the number of fields.
func (r *Reader) FieldCount() int {
	return len(r.Fields())
}

// ShapeType returns the shape type that was declared in the file header of the
// .shp structure.
func (r *Reader) ShapeType() shp.ShapeType {
	return r.shp.ShapeType
}

// BBox returns the bounding box that was declared in the file header of the
// .shp structure.
func (r *Reader) BBox() *shp.BBox {
	return &r.shp.BBox
}

// Next reads the next record from the underlying readers. When the end of the
// shp reader is reached, this will return io.EOF. If a dbf reader is provided
// that has fewer records than the given shp reader, io.ErrUnexpectedEOF will
// be returned after the last record in the db reader.
func (r *Reader) Next() (*Record, error) {
	s, err := r.shp.Next()
	if err != nil {
		return nil, err
	}

	if r.dbf == nil {
		return &Record{Shape: s}, nil
	}

	a, err := r.dbf.Next()
	if err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	} else if err != nil {
		return nil, err
	}

	return &Record{
		Shape:  s,
		Record: a,
	}, nil
}

// NewReader creates a new Reader that provides sequential access to the
// records in a shp file structure. Optionally, a dbf io.Reader can be
// provided that will also provide access to the attributes that correspond
// to each shp record.
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
