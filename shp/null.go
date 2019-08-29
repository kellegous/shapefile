package shp

import (
	"fmt"
	"io"
)

// Null indicates a shape with no geometric data. Each feature type (point,
// line, polygon, etc.) supports nulls. It is valid to have points and null
// points in the same shapefile.
type Null struct{}

// Type ...
func (n Null) Type() ShapeType {
	return TypeNull
}

func readNull(r io.Reader, cl int32) (*Null, error) {
	if cl != 2 {
		return nil, fmt.Errorf("unexpected content length for null: %d", cl)
	}
	return &Null{}, nil
}
