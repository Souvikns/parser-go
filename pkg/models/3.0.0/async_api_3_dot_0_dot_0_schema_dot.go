
package models

// AsyncApi_3Dot_0Dot_0SchemaDot represents a AsyncApi_3Dot_0Dot_0SchemaDot model.
type AsyncApi_3Dot_0Dot_0SchemaDot struct {
  Asyncapi string `json:"asyncapi,omitempty"`
  Id string `json:"id,omitempty"`
  Info *Info `json:"info,omitempty"`
  Servers map[string]ServersAdditionalProperty `json:"servers,omitempty"`
  DefaultContentType string `json:"defaultContentType,omitempty"`
  Channels map[string]ChannelsAdditionalProperty `json:"channels,omitempty"`
  Operations map[string]OperationsAdditionalProperty `json:"operations,omitempty"`
  Components *Components `json:"components,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}