
package models

// AsyncApi_2Dot_0Dot_0MinusRc1SchemaDot represents a AsyncApi_2Dot_0Dot_0MinusRc1SchemaDot model.
type AsyncApi_2Dot_0Dot_0MinusRc1SchemaDot struct {
  Asyncapi *Asyncapi
  Id string
  Info *Info
  Servers []Server
  DefaultContentType string
  Channels map[string]ChannelItem
  Components *Components
  Tags []Tag
  ExternalDocs *ExternalDocs
  AdditionalProperties map[string]interface{}
}