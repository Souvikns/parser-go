
package models

// BearerHttpSecurityScheme represents a BearerHttpSecurityScheme model.
type BearerHttpSecurityScheme struct {
  Scheme *BearerHttpSecuritySchemeScheme `json:"scheme,omitempty"`
  BearerFormat string `json:"bearerFormat,omitempty"`
  ReservedType *BearerHttpSecuritySchemeType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}