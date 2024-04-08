
package models

// AvroTypes represents a AvroTypes model.
type AvroTypes struct {
  PrimitiveType
  PrimitiveTypeWithMetadata
  ModelinaAnyType interface{}
  Record
  Enum
  Array
  ReservedMap
  Fixed
  ModelinaArrType []AvroSchema
}