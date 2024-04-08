
package models

// SecurityScheme represents a SecurityScheme model.
type SecurityScheme struct {
  UserPassword
  ApiKey
  X509
  SymmetricEncryption
  AsymmetricEncryption
  HttpSecurityScheme
}