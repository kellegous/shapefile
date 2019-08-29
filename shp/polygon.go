package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Polygon consists of one or more rings. A ring is a connected sequence of
// four or more points that form a closed, non-self-intersecting loop. A
// polygon may contain multiple outer rings. The order of vertices or
// orientation for a ring indicates which side of the ring is the interior
// of the polygon.  The neighborhood to the right of an observer walking
// along the ring in vertex order is the neighborhood inside the polygon.
// Vertices of rings defining holes in polygons are in a counterclockwise
// direction.  Vertices for a single, ringed polygon are, therefore, always
// in clockwise order.  The rings of a polygon are referred to as its parts.
type Polygon struct {
	BBox           BBox
	NumberOfParts  int32
	NumberOfPoints int32
	Parts          []int32
	Points         []Point
}

// Type ...
func (p *Polygon) Type() ShapeType {
	return TypePolygon
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
