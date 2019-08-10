package shp

import (
	"encoding/binary"
	"io"
	"math"
)

// MData encapsulates all the optional data for M type
// shape records.
type MData struct {
	MRange Range
	M      []float64
}

func (m *MData) empty(n int32) {
	nan := math.NaN()
	m.MRange.Min = nan
	m.MRange.Max = nan

	m.M = make([]float64, n)
	for i := int32(0); i < n; i++ {
		m.M[i] = nan
	}
}

func (m *MData) read(r io.Reader, n int32) error {
	// Range
	if err := binary.Read(r, binary.LittleEndian, &m.MRange); err != nil {
		return err
	}

	// M
	m.M = make([]float64, n)
	if err := binary.Read(r, binary.LittleEndian, &m.M); err != nil {
		return err
	}

	m.MRange.Min = doubleToFloat64(m.MRange.Min)
	m.MRange.Max = doubleToFloat64(m.MRange.Max)
	for i, n := 0, len(m.M); i < n; i++ {
		m.M[i] = doubleToFloat64(m.M[i])
	}

	return nil
}
