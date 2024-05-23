
package models

// SaslScramSecurityScheme represents a SaslScramSecurityScheme model.
type SaslScramSecurityScheme struct {
  ReservedType *SaslScramSecuritySchemeType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}