
package models

// OpenIdConnect represents a OpenIdConnect model.
type OpenIdConnect struct {
  ReservedType *OpenIdConnectType `json:"type"`
  Description string `json:"description"`
  OpenIdConnectUrl string `json:"openIdConnectUrl"`
  Scopes []string `json:"scopes"`
  AdditionalProperties map[string]interface{} `json:"-"`
}