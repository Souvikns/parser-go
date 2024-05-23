
package models

// SaslGssapiSecurityScheme represents a SaslGssapiSecurityScheme model.
type SaslGssapiSecurityScheme struct {
  ReservedType *SaslGssapiSecuritySchemeType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}