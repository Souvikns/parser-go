
package models

// SaslPlainSecurityScheme represents a SaslPlainSecurityScheme model.
type SaslPlainSecurityScheme struct {
  ReservedType *SaslPlainSecuritySchemeType `json:type`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}