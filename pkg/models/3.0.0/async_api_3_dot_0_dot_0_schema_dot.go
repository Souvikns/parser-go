
package models

// AsyncApi_3Dot_0Dot_0SchemaDot represents a AsyncApi_3Dot_0Dot_0SchemaDot model.
type AsyncApi_3Dot_0Dot_0SchemaDot struct {
  Asyncapi string `json:"asyncapi"`
  Id string `json:"id"`
  Info *Info `json:"info"`
  Servers map[string]ServersAdditionalProperty `json:"servers"`
  DefaultContentType string `json:"defaultContentType"`
  Channels map[string]ChannelsAdditionalProperty `json:"channels"`
  Operations map[string]OperationsAdditionalProperty `json:"operations"`
  Components *Components `json:"components"`
  AdditionalProperties map[string]interface{} `json:"-"`
}