
package models

// Oauth2FlowsFlowsAuthorizationCode represents a Oauth2FlowsFlowsAuthorizationCode model.
type Oauth2FlowsFlowsAuthorizationCode struct {
  AuthorizationUrl string `json:"authorizationUrl"`
  TokenUrl string `json:"tokenUrl"`
  RefreshUrl string `json:"refreshUrl"`
  AvailableScopes map[string]string `json:"availableScopes"`
  AdditionalProperties map[string]interface{} `json:"-"`
}