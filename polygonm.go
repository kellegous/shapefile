package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PolygonM ...
type PolygonM struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
	Range          Range
	Measures       []float64
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
	p.Measures = make([]float64, p.NumberOfPoints)
	if cl == min {
		initNanMeasures(&p.Range, p.Measures)
	} else if err := readMeasures(r, &p.Range, p.Measures); err != nil {
		return nil, err
	}

	return &p, nil
}
