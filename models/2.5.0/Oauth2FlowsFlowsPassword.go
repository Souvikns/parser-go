
package models

// Oauth2FlowsFlowsPassword represents a Oauth2FlowsFlowsPassword model.
type Oauth2FlowsFlowsPassword struct {
  AuthorizationUrl string `json:authorizationUrl`
  TokenUrl string `json:tokenUrl`
  RefreshUrl string `json:refreshUrl`
  Scopes map[string]string `json:scopes`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}