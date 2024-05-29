
package models

// OpenIdConnect represents a OpenIdConnect model.
type OpenIdConnect struct {
  ReservedType *OpenIdConnectType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  OpenIdConnectUrl string `json:"openIdConnectUrl,omitempty"`
  Scopes []string `json:"scopes,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}