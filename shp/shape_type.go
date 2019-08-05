package shp

// ShapeType ...
type ShapeType int32

const (
	// TypeNull ...
	TypeNull = 0

	// TypePoint ...
	TypePoint = 1

	// TypePolyline ...
	TypePolyline = 3

	// TypePolygon ...
	TypePolygon = 5

	// TypeMultiPoint ...
	TypeMultiPoint = 8

	// TypePointZ ...
	TypePointZ = 11

	// TypePolylineZ ...
	TypePolylineZ = 13

	// TypePolygonZ ...
	TypePolygonZ = 15

	// TypeMultiPointZ ...
	TypeMultiPointZ = 18

	// TypePointM ...
	TypePointM = 21

	// TypePolylineM ...
	TypePolylineM = 23

	// TypePolygonM ...
	TypePolygonM = 25

	// TypeMultiPointM ...
	TypeMultiPointM = 28

	// TypeMultiPatch ...
	TypeMultiPatch = 31
)

func (s ShapeType) String() string {
	switch s {
	case TypeNull:
		return "Null"
	case TypePoint:
		return "Point"
	case TypePolyline:
		return "Polyline"
	case TypePolygon:
		return "Polygon"
	case TypeMultiPoint:
		return "MultiPoint"
	case TypePointZ:
		return "PointZ"
	case TypePolylineZ:
		return "PolylineZ"
	case TypePolygonZ:
		return "PolygonZ"
	case TypeMultiPointZ:
		return "MultiPointZ"
	case TypePointM:
		return "PointM"
	case TypePolylineM:
		return "PolylineM"
	case TypePolygonM:
		return "PolygonM"
	case TypeMultiPatch:
		return "MultiPatch"
	}
	return "Unknown"
}
