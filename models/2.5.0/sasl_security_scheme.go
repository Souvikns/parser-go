
package models

// SaslSecurityScheme represents a SaslSecurityScheme model.
type SaslSecurityScheme struct {
  SaslPlainSecurityScheme
  SaslScramSecurityScheme
  SaslGssapiSecurityScheme
}