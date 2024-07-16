
package models

// SecurityScheme represents a SecurityScheme model.
type SecurityScheme struct {
  UserPassword `json:"-,omitempty`
  ApiKey `json:"-,omitempty`
  X509 `json:"-,omitempty`
  SymmetricEncryption `json:"-,omitempty`
  AsymmetricEncryption `json:"-,omitempty`
  HttpSecurityScheme `json:"-,omitempty`
  Oauth2Flows `json:"-,omitempty`
  OpenIdConnect `json:"-,omitempty`
  SaslSecurityScheme `json:"-,omitempty`
}