
package models

// Oauth2Flows represents a Oauth2Flows model.
type Oauth2Flows struct {
  ReservedType *Oauth2FlowsType `json:type`
  Description string `json:description`
  Flows *Oauth2FlowsFlows `json:flows`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}