
package models

// AsyncApi_2Dot_0Dot_0MinusRc1SchemaDot represents a AsyncApi_2Dot_0Dot_0MinusRc1SchemaDot model.
type AsyncApi_2Dot_0Dot_0MinusRc1SchemaDot struct {
  Asyncapi *Asyncapi `json:"asyncapi,omitempty"`
  Id string `json:"id,omitempty"`
  Info *Info `json:"info,omitempty"`
  Servers []Server `json:"servers,omitempty"`
  DefaultContentType string `json:"defaultContentType,omitempty"`
  Channels map[string]ChannelItem `json:"channels,omitempty"`
  Components *Components `json:"components,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}