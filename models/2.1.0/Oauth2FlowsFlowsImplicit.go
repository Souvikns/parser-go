
package models

// Oauth2FlowsFlowsImplicit represents a Oauth2FlowsFlowsImplicit model.
type Oauth2FlowsFlowsImplicit struct {
  AuthorizationUrl string `json:authorizationUrl`
  TokenUrl string `json:tokenUrl`
  RefreshUrl string `json:refreshUrl`
  Scopes map[string]string `json:scopes`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}