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

	// TypeMultipoint ...
	TypeMultipoint = 8

	// TypePointZ ...
	TypePointZ = 11

	// TypePolylineZ ...
	TypePolylineZ = 13

	// TypePolygonZ ...
	TypePolygonZ = 15

	// TypeMultipointZ ...
	TypeMultipointZ = 18

	// TypePointM ...
	TypePointM = 21

	// TypeLineM ...
	TypeLineM = 23

	// TypePolylineM ...
	TypePolylineM = 23

	// TypePolygonM ...
	TypePolygonM = 25

	// TypeMultiPointM ...
	TypeMultiPointM = 28

	// TypeMultiPatch ...
	TypeMultiPatch = 31
)
