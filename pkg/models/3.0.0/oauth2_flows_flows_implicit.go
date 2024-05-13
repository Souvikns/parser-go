
package models

// Oauth2FlowsFlowsImplicit represents a Oauth2FlowsFlowsImplicit model.
type Oauth2FlowsFlowsImplicit struct {
  AuthorizationUrl string `json:"authorizationUrl"`
  TokenUrl string `json:"tokenUrl"`
  RefreshUrl string `json:"refreshUrl"`
  AvailableScopes map[string]string `json:"availableScopes"`
  AdditionalProperties map[string]interface{} `json:"-"`
}