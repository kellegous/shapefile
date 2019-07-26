package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Point ...
type Point struct {
	X, Y float64
}

func readPoint(r io.Reader, cl int32) (*Point, error) {
	if cl != 10 {
		return nil, fmt.Errorf("unexpected content length for Point: %d", cl)
	}

	var pt Point
	if err := binary.Read(r, binary.LittleEndian, &pt); err != nil {
		return nil, err
	}

	return &pt, nil
}
