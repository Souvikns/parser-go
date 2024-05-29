
package models

// Oauth2FlowsFlowsAuthorizationCode represents a Oauth2FlowsFlowsAuthorizationCode model.
type Oauth2FlowsFlowsAuthorizationCode struct {
  AuthorizationUrl string `json:"authorizationUrl,omitempty"`
  TokenUrl string `json:"tokenUrl,omitempty"`
  RefreshUrl string `json:"refreshUrl,omitempty"`
  Scopes map[string]string `json:"scopes,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}