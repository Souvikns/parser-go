
package models

// Oauth2FlowsFlows represents a Oauth2FlowsFlows model.
type Oauth2FlowsFlows struct {
  Implicit *Oauth2FlowsFlowsImplicit `json:"implicit,omitempty"`
  Password *Oauth2FlowsFlowsPassword `json:"password,omitempty"`
  ClientCredentials *Oauth2FlowsFlowsClientCredentials `json:"clientCredentials,omitempty"`
  AuthorizationCode *Oauth2FlowsFlowsAuthorizationCode `json:"authorizationCode,omitempty"`
}