package shp

// ShapeType is an enum that specifies the type of the associated Shape.
type ShapeType int32

const (
	// TypeNull indicates Null
	TypeNull = 0

	// TypePoint indicates Point
	TypePoint = 1

	// TypePolyline indicates Polyline
	TypePolyline = 3

	// TypePolygon indicates Polygon
	TypePolygon = 5

	// TypeMultiPoint indicates MultiPoint
	TypeMultiPoint = 8

	// TypePointZ indicates PointZ
	TypePointZ = 11

	// TypePolylineZ indicates PolylineZ
	TypePolylineZ = 13

	// TypePolygonZ indicates PolygonZ
	TypePolygonZ = 15

	// TypeMultiPointZ indicates MultiPointZ
	TypeMultiPointZ = 18

	// TypePointM indicates PointM
	TypePointM = 21

	// TypePolylineM indicates PolylineM
	TypePolylineM = 23

	// TypePolygonM indicates PolygonM
	TypePolygonM = 25

	// TypeMultiPointM indicates MultiPointM
	TypeMultiPointM = 28

	// TypeMultiPatch indicates MultiPatch
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
