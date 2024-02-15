
package models

// OpenIdConnect represents a OpenIdConnect model.
type OpenIdConnect struct {
  ReservedType *OpenIdConnectType
  Description string
  OpenIdConnectUrl string
  Scopes []string
  AdditionalProperties map[string]interface{}
}