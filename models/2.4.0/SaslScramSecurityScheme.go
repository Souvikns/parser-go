
package models

// SaslScramSecurityScheme represents a SaslScramSecurityScheme model.
type SaslScramSecurityScheme struct {
  ReservedType *SaslScramSecuritySchemeType `json:type`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}