
package models

// PrimitiveType represents an enum of PrimitiveType.
type PrimitiveType uint

const (
  PrimitiveTypeNull PrimitiveType = iota
  PrimitiveTypeBoolean
  PrimitiveTypeInt
  PrimitiveTypeLong
  PrimitiveTypeFloat
  PrimitiveTypeDouble
  PrimitiveTypeBytes
  PrimitiveTypeString
)

// Value returns the value of the enum.
func (op PrimitiveType) Value() any {
	if op >= PrimitiveType(len(PrimitiveTypeValues)) {
		return nil
	}
	return PrimitiveTypeValues[op]
}

var PrimitiveTypeValues = []any{"null","boolean","int","long","float","double","bytes","string"}
var ValuesToPrimitiveType = map[any]PrimitiveType{
  PrimitiveTypeValues[PrimitiveTypeNull]: PrimitiveTypeNull,
  PrimitiveTypeValues[PrimitiveTypeBoolean]: PrimitiveTypeBoolean,
  PrimitiveTypeValues[PrimitiveTypeInt]: PrimitiveTypeInt,
  PrimitiveTypeValues[PrimitiveTypeLong]: PrimitiveTypeLong,
  PrimitiveTypeValues[PrimitiveTypeFloat]: PrimitiveTypeFloat,
  PrimitiveTypeValues[PrimitiveTypeDouble]: PrimitiveTypeDouble,
  PrimitiveTypeValues[PrimitiveTypeBytes]: PrimitiveTypeBytes,
  PrimitiveTypeValues[PrimitiveTypeString]: PrimitiveTypeString,
}
