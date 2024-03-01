
package models

// Oauth2FlowsFlows represents a Oauth2FlowsFlows model.
type Oauth2FlowsFlows struct {
  Implicit *Oauth2FlowsFlowsImplicit `json:"implicit"`
  Password *Oauth2FlowsFlowsPassword `json:"password"`
  ClientCredentials *Oauth2FlowsFlowsClientCredentials `json:"clientCredentials"`
  AuthorizationCode *Oauth2FlowsFlowsAuthorizationCode `json:"authorizationCode"`
}