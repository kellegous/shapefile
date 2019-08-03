package shp

import (
	"math"
)

func doubleToFloat64(v float64) float64 {
	if v < -1e38 {
		return math.NaN()
	}
	return v
}
