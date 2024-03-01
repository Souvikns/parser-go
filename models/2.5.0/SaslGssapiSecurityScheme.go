
package models

// SaslGssapiSecurityScheme represents a SaslGssapiSecurityScheme model.
type SaslGssapiSecurityScheme struct {
  ReservedType *SaslGssapiSecuritySchemeType `json:type`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}