package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Polygon ...
type Polygon struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
}

func readPolygon(r io.Reader, cl int32) (*Polygon, error) {
	var p Polygon

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &p.BBox); err != nil {
		return nil, err
	}

	// NumberOfParts
	if err := binary.Read(r, binary.LittleEndian, &p.NumberOfParts); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &p.NumberOfPoints); err != nil {
		return nil, err
	}

	if cl != 22+p.NumberOfPoints*8+p.NumberOfParts*2 {
		return nil, fmt.Errorf("invalid content length for Polygon: %d", cl)
	}

	// Parts
	p.Parts = make([]int32, p.NumberOfParts)
	if err := binary.Read(r, binary.LittleEndian, &p.Parts); err != nil {
		return nil, err
	}

	// Points
	p.Points = make([]Point, p.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &p.Points); err != nil {
		return nil, err
	}

	return &p, nil
}
