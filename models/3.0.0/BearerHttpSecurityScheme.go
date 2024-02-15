
package models

// BearerHttpSecurityScheme represents a BearerHttpSecurityScheme model.
type BearerHttpSecurityScheme struct {
  Scheme *BearerHttpSecuritySchemeScheme
  BearerFormat string
  ReservedType *BearerHttpSecuritySchemeType
  Description string
  AdditionalProperties map[string]interface{}
}