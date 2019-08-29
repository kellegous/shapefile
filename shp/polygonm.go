package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PolygonM is a Polygon with optional measure (M) data. Missing M data is
// specified as NaN.
type PolygonM struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
	MData
}

// Type ...
func (p *PolygonM) Type() ShapeType {
	return TypePointM
}

func readPolygonM(r io.Reader, cl int32) (*PolygonM, error) {
	var p PolygonM

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

	min := 22 + 2*p.NumberOfParts + 8*p.NumberOfPoints
	max := min + 8 + 4*p.NumberOfPoints

	if cl != min && cl != max {
		return nil, fmt.Errorf("invalid content length for PolygonM: %d", cl)
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

	// Measures
	if cl == min {
		p.MData.empty(p.NumberOfPoints)
	} else if err := p.MData.read(r, p.NumberOfPoints); err != nil {
		return nil, err
	}

	return &p, nil
}
