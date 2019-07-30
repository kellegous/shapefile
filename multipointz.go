package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MultiPointZ ...
type MultiPointZ struct {
	BBox           BBox
	NumberOfPoints int32
	Points         []Point
	ZRange         Range
	Z              []float64
	M
}

func readMultiPointZ(r io.Reader, cl int32) (*MultiPointZ, error) {
	var s MultiPointZ

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &s.BBox); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &s.NumberOfPoints); err != nil {
		return nil, err
	}

	min := 28 + 12*s.NumberOfPoints
	max := min + 8 + 4*s.NumberOfPoints
	if cl != min && cl != max {
		return nil, fmt.Errorf("invalid content length for MultiPointZ: %d", cl)
	}

	// Points
	s.Points = make([]Point, s.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &s.Points); err != nil {
		return nil, err
	}

	// ZRange
	if err := binary.Read(r, binary.LittleEndian, &s.ZRange); err != nil {
		return nil, err
	}

	// Z
	s.Z = make([]float64, s.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &s.Z); err != nil {
		return nil, err
	}

	// Measures
	if cl == min {
		s.M.empty(s.NumberOfPoints)
	} else if err := s.M.read(r, s.NumberOfPoints); err != nil {
		return nil, err
	}

	return &s, nil
}
