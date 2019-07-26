package shp

import (
	"fmt"
	"io"
)

// Null ...
type Null struct{}

func readNull(r io.Reader, cl int32) (*Null, error) {
	if cl != 2 {
		return nil, fmt.Errorf("unexpected content length for null: %d", cl)
	}
	return &Null{}, nil
}
