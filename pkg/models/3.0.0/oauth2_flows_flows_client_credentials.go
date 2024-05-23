
package models

// Oauth2FlowsFlowsClientCredentials represents a Oauth2FlowsFlowsClientCredentials model.
type Oauth2FlowsFlowsClientCredentials struct {
  AuthorizationUrl string `json:"authorizationUrl,omitempty"`
  TokenUrl string `json:"tokenUrl,omitempty"`
  RefreshUrl string `json:"refreshUrl,omitempty"`
  AvailableScopes map[string]string `json:"availableScopes,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}