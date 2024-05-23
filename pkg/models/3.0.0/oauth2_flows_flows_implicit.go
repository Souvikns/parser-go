
package models

// Oauth2FlowsFlowsImplicit represents a Oauth2FlowsFlowsImplicit model.
type Oauth2FlowsFlowsImplicit struct {
  AuthorizationUrl string `json:"authorizationUrl,omitempty"`
  TokenUrl string `json:"tokenUrl,omitempty"`
  RefreshUrl string `json:"refreshUrl,omitempty"`
  AvailableScopes map[string]string `json:"availableScopes,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}