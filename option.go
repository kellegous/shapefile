package shapefile

import (
	"io"

	"github.com/kellegous/shapefile/dbf"
)

// Option allows for options to be specified when calling NewReader.
type Option func(r *Reader) error

// WithDBF allows the specification of a dbf file that
// contains the associated attributes for each of the shp
// file records.
func WithDBF(r io.Reader) Option {
	return func(rr *Reader) error {
		dr, err := dbf.NewReader(r)
		if err != nil {
			return err
		}
		rr.dbf = dr
		return nil
	}
}
