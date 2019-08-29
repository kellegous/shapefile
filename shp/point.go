package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Point consists of a pair of double-precision coordinates in the order X,Y.
type Point struct {
	X, Y float64
}

// Type ...
func (p *Point) Type() ShapeType {
	return TypePoint
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
