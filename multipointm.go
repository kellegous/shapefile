package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// MultiPointM ...
type MultiPointM struct {
	BBox           BBox
	NumberOfPoints int32
	Points         []Point
	Range          Range
	Measures       []float64
}

func sizeOfMultiPointM(n int32) int32 {
	return 28 + n*8 + n*4
}

func readMultiPointM(r io.Reader, cl int32) (*MultiPointM, error) {
	var mp MultiPointM

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &mp.BBox); err != nil {
		return nil, err
	}

	// NumberOfPoints
	if err := binary.Read(r, binary.LittleEndian, &mp.NumberOfPoints); err != nil {
		return nil, err
	}

	min := 20 + mp.NumberOfPoints*8
	max := min + 8 + mp.NumberOfPoints*4

	if cl != min && cl != max {
		return nil, fmt.Errorf("invalid content length for MultiPointM: %d", cl)
	}

	// Points
	mp.Points = make([]Point, mp.NumberOfPoints)
	if err := binary.Read(r, binary.LittleEndian, &mp.Points); err != nil {
		return nil, err
	}

	// Measures
	mp.Measures = make([]float64, mp.NumberOfPoints)
	if cl == min {
		initNanMeasures(&mp.Range, mp.Measures)
	} else if err := readMeasures(r, &mp.Range, mp.Measures); err != nil {
		return nil, err
	}

	return &mp, nil
}
