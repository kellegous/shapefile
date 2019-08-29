package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PartType ...
type PartType int32

const (
	// TriangleStrip indicates the part is a linked strip of triangles,
	// where every vertex (after the first two) completes a new triangle.
	// A new triangle is always formed byconnecting the new vertex with
	// its two immediate predecessors.
	TriangleStrip PartType = 0

	// TriangleFan indicates the part is a linked fan of triangles,
	// where every vertex (after the first two) completes a new triangle.
	// A new triangle is always formed byconnecting the new vertex with
	// its immediate predecessor and thefirst vertex of the part
	TriangleFan PartType = 1

	// OuterRing indicates the part is the outer ring of a polygon.
	OuterRing PartType = 2

	// InnerRing indicates the part is a hole in a polygon.
	InnerRing PartType = 3

	// FirstRing indicates the part is the first ring of a polygon of
	// unspecified type.
	FirstRing PartType = 4

	// Ring indicates the part is a ring of a polygon of unspecified type.
	Ring PartType = 5
)

// MultiPatch consists of a number of surface patches.  Each surface patch
// describes a surface.  The surface patches of a MultiPatch are referred
// to as its parts, and the type ofpart controls how the order of vertices
// of an MultiPatch part is interpreted.
type MultiPatch struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	PartTypes      []PartType
	Points         []Point
	ZRange         Range
	Z              []float64
	MData
}

// Type ...
func (m *MultiPatch) Type() ShapeType {
	return TypeMultiPatch
}

func readMultiPatch(r io.Reader, cl int32) (*MultiPatch, error) {
	var s MultiPatch

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

	min := 30 + 4*s.NumberOfParts + 12*s.NumberOfPoints
	max := min + 8 + 4*s.NumberOfPoints
	if cl != min && cl != max {
		return nil, fmt.Errorf("invalid content length for MultiPatch: %d", cl)
	}

	// Parts
	s.Parts = make([]int32, s.NumberOfParts)
	if err := binary.Read(r, binary.LittleEndian, &s.Parts); err != nil {
		return nil, err
	}

	// PartTypes
	s.PartTypes = make([]PartType, s.NumberOfParts)
	if err := binary.Read(r, binary.LittleEndian, &s.PartTypes); err != nil {
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
