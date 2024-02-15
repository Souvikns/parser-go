
package models

// Oauth2Flows represents a Oauth2Flows model.
type Oauth2Flows struct {
  ReservedType *Oauth2FlowsType
  Description string
  Flows *Oauth2FlowsFlows
  Scopes []string
  AdditionalProperties map[string]interface{}
}