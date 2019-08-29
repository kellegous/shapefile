package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PolylineZ represents a polyline in the order X,Y,Z with optional
// associated measure (M) data. Missing M data is indicated with NaN.
type PolylineZ struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
	ZRange         Range
	Z              []float64
	MData
}

// Type ...
func (p *PolylineZ) Type() ShapeType {
	return TypePolylineZ
}

func readPolylineZ(r io.Reader, cl int32) (*PolylineZ, error) {
	var s PolylineZ

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &s.BBox); err != nil {
		return nil, err
	}

	// NumberOfParts
	if err := binary.Read(r, binary.LittleEndian, &s.NumberOfParts); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &s.NumberOfPoints); err != nil {
		return nil, err
	}

	min := 30 + 2*s.NumberOfParts + 12*s.NumberOfPoints
	max := min + 8 + 4*s.NumberOfPoints
	if cl != min && cl != max {
		return nil, fmt.Errorf("invalid content length for PolylineZ: %d", cl)
	}

	// Parts
	s.Parts = make([]int32, s.NumberOfParts)
	if err := binary.Read(r, binary.LittleEndian, &s.Parts); err != nil {
		return nil, err
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

	// MData
	if cl == min {
		s.MData.empty(s.NumberOfPoints)
	} else if err := s.MData.read(r, s.NumberOfPoints); err != nil {
		return nil, err
	}

	return &s, nil
}
