
package models

// SaslSecurityScheme represents a SaslSecurityScheme model.
type SaslSecurityScheme struct {
  SaslPlainSecurityScheme `json:"-,omitempty`
  SaslScramSecurityScheme `json:"-,omitempty`
  SaslGssapiSecurityScheme `json:"-,omitempty`
}