
package models

// SaslPlainSecurityScheme represents a SaslPlainSecurityScheme model.
type SaslPlainSecurityScheme struct {
  ReservedType *SaslPlainSecuritySchemeType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}