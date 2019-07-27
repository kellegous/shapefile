package shp

import (
	"encoding/binary"
	"io"
	"math"
)

// M ...
type M struct {
	Range    Range
	Measures []float64
}

func (m *M) empty(n int32) {
	nan := math.NaN()
	m.Range.Min = nan
	m.Range.Max = nan

	m.Measures = make([]float64, n)
	for i := int32(0); i < n; i++ {
		m.Measures[i] = nan
	}
}

func (m *M) read(r io.Reader, n int32) error {
	// Range
	if err := binary.Read(r, binary.LittleEndian, &m.Range); err != nil {
		return err
	}

	// Measures
	m.Measures = make([]float64, n)
	if err := binary.Read(r, binary.LittleEndian, &m.Measures); err != nil {
		return err
	}

	m.Range.Min = doubleToFloat64(m.Range.Min)
	m.Range.Max = doubleToFloat64(m.Range.Max)
	for i, n := 0, len(m.Measures); i < n; i++ {
		m.Measures[i] = doubleToFloat64(m.Measures[i])
	}

	return nil
}
