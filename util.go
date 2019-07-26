package shp

import (
	"encoding/binary"
	"io"
	"math"
)

func doubleToFloat64(v float64) float64 {
	if v < -1e38 {
		return math.NaN()
	}
	return v
}

func readMeasures(r io.Reader, rng *Range, m []float64) error {
	// Range ...
	if err := binary.Read(r, binary.LittleEndian, rng); err != nil {
		return err
	}

	// Measures ...
	if err := binary.Read(r, binary.LittleEndian, m); err != nil {
		return err
	}

	// convert any nodata entries to NaN
	rng.Min = doubleToFloat64(rng.Min)
	rng.Max = doubleToFloat64(rng.Max)
	for i, n := 0, len(m); i < n; i++ {
		m[i] = doubleToFloat64(m[i])
	}

	return nil
}

func initNanMeasures(rng *Range, m []float64) {
	nan := math.NaN()
	rng.Min = nan
	rng.Max = nan
	for i, n := 0, len(m); i < n; i++ {
		m[i] = nan
	}
}
