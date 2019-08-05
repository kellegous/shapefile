package shapefile

import (
	"github.com/kellegous/shapefile/dbf"
	"github.com/kellegous/shapefile/shp"
)

// Record ...
type Record struct {
	Shape shp.Shape
	*dbf.Record
}

// Attrs ...
func (r *Record) Attrs() []string {
	if r.Record == nil {
		return nil
	}
	return r.Record.Attrs()
}
