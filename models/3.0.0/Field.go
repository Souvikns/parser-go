
package models

// Field represents a Field model.
type Field struct {
  Name string
  ReservedType interface{}
  Doc string
  ReservedDefault interface{}
  Order *AvroSchemaV1AvroFieldOrder
  Aliases []string
  AdditionalProperties map[string]interface{}
}