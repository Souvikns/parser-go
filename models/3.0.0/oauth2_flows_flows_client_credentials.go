
package models

// Oauth2FlowsFlowsClientCredentials represents a Oauth2FlowsFlowsClientCredentials model.
type Oauth2FlowsFlowsClientCredentials struct {
  AuthorizationUrl string `json:"authorizationUrl"`
  TokenUrl string `json:"tokenUrl"`
  RefreshUrl string `json:"refreshUrl"`
  AvailableScopes map[string]string `json:"availableScopes"`
  AdditionalProperties map[string]interface{} `json:"-"`
}