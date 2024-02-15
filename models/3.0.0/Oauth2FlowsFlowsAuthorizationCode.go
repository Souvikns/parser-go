
package models

// Oauth2FlowsFlowsAuthorizationCode represents a Oauth2FlowsFlowsAuthorizationCode model.
type Oauth2FlowsFlowsAuthorizationCode struct {
  AuthorizationUrl string
  TokenUrl string
  RefreshUrl string
  AvailableScopes map[string]string
  AdditionalProperties map[string]interface{}
}