package shp

import "io"

// Option ...
type Option func(r *Reader)

// WithDBF ...
func WithDBF(dbf io.Reader) Option {
	return func(r *Reader) {
		r.dbf = dbf
	}
}
