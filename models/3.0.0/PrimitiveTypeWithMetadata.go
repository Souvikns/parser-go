
package models

// PrimitiveTypeWithMetadata represents a PrimitiveTypeWithMetadata model.
type PrimitiveTypeWithMetadata struct {
  ReservedType *PrimitiveType `json:"type"`
  AdditionalProperties map[string]interface{} `json:"-"`
}