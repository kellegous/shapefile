package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PointZ is a triplet of double-precision coordinates in the order X,Y,Z, plus
// a measure.
type PointZ struct {
	X, Y, Z, M float64
}

// Type ...
func (p *PointZ) Type() ShapeType {
	return TypePointZ
}

func readPointZ(r io.Reader, cl int32) (*PointZ, error) {
	if cl != 18 {
		return nil, fmt.Errorf("unexpected content length for PointZ: %d", cl)
	}

	var p PointZ
	if err := binary.Read(r, binary.LittleEndian, &p); err != nil {
		return nil, err
	}

	p.M = doubleToFloat64(p.M)

	return &p, nil
}
