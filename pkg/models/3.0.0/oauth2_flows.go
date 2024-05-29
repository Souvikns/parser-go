
package models

// Oauth2Flows represents a Oauth2Flows model.
type Oauth2Flows struct {
  ReservedType *Oauth2FlowsType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  Flows *Oauth2FlowsFlows `json:"flows,omitempty"`
  Scopes []string `json:"scopes,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}