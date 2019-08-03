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
