package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Polyline is an ordered set of vertices that consists of one or more parts.
// A part is a connected sequence of two or more points.  Parts may or may not
// be connected to one another. Parts may or may not intersect one another.
type Polyline struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
}

// Type ...
func (p *Polyline) Type() ShapeType {
	return TypePolyline
}

func readPolyline(r io.Reader, cl int32) (*Polyline, error) {
	var pl Polyline
	if err := binary.Read(r, binary.LittleEndian, &pl.BBox); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &pl.NumberOfParts); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &pl.NumberOfPoints); err != nil {
		return nil, err
	}

	if cl != 22+pl.NumberOfPoints*8+pl.NumberOfParts*2 {
		return nil, fmt.Errorf("invalid content length for Polyline: %d", cl)
	}

	pl.Parts = make([]int32, pl.NumberOfParts)
	for i := int32(0); i < pl.NumberOfParts; i++ {
		if err := binary.Read(r, binary.LittleEndian, &pl.Parts[i]); err != nil {
			return nil, err
		}
	}

	pl.Points = make([]Point, pl.NumberOfPoints)
	for i := int32(0); i < pl.NumberOfPoints; i++ {
		if err := binary.Read(r, binary.LittleEndian, &pl.Points[i]); err != nil {
			return nil, err
		}
	}

	return &pl, nil
}
