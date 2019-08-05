package shp

import (
	"encoding/binary"
	"fmt"
	"io"
)

// PointM ...
type PointM struct {
	X, Y, M float64
}

func readPointM(r io.Reader, cl int32) (*PointM, error) {
	if cl != 14 {
		return nil, fmt.Errorf("unexpected content length for PointM: %d", cl)
	}

	var pt PointM
	if err := binary.Read(r, binary.LittleEndian, &pt); err != nil {
		return nil, err
	}

	pt.M = doubleToFloat64(pt.M)

	return &pt, nil
}