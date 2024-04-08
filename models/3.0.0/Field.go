
package models

// Field represents a Field model.
type Field struct {
  Name string `json:"name"`
  ReservedType *AvroTypes `json:"type"`
  Doc string `json:"doc"`
  ReservedDefault interface{} `json:"default"`
  Order *AvroSchemaV1AvroFieldOrder `json:"order"`
  Aliases []string `json:"aliases"`
  AdditionalProperties map[string]interface{} `json:"-"`
}