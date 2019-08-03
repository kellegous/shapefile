package dbf

// FieldType ...
type FieldType byte

const (
	// TypeCharacter ...
	TypeCharacter FieldType = 'C'

	// TypeDate ...
	TypeDate FieldType = 'D'

	// TypeFloat ...
	TypeFloat FieldType = 'F'

	// TypeLogical ...
	TypeLogical FieldType = 'L'

	// TypeNumeric ...
	TypeNumeric FieldType = 'N'
)

// Field ...
type Field struct {
	Name         string
	Type         FieldType
	Length       byte
	DecimalCount byte
	off          int
}
