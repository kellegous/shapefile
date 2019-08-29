package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MultiPoint represents a set of points.
type MultiPoint struct {
	BBox           BBox
	NumberOfPoints int32
	Points         []Point
}

// Type ...
func (m *MultiPoint) Type() ShapeType {
	return TypeMultiPoint
}

func readMultiPoint(r io.Reader, cl int32) (*MultiPoint, error) {
	var mp MultiPoint

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &mp.BBox); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &mp.NumberOfPoints); err != nil {
		return nil, err
	}

	if cl != 20+mp.NumberOfPoints*8 {
		return nil, fmt.Errorf("invalid content length for MultiPoint: %d", cl)
	}

	// Points
	mp.Points = make([]Point, mp.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &mp.Points); err != nil {
		return nil, err
	}

	return &mp, nil
}
