package shapefile

import (
	"io"

	"github.com/kellegous/shapefile/dbf"
)

// Option ...
type Option func(r *Reader) error

// WithDBF ...
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
