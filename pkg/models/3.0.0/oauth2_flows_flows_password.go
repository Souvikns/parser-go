
package models

// Oauth2FlowsFlowsPassword represents a Oauth2FlowsFlowsPassword model.
type Oauth2FlowsFlowsPassword struct {
  AuthorizationUrl string `json:"authorizationUrl,omitempty"`
  TokenUrl string `json:"tokenUrl,omitempty"`
  RefreshUrl string `json:"refreshUrl,omitempty"`
  AvailableScopes map[string]string `json:"availableScopes,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}