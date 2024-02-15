
package models

// AsyncApi_2Dot_0Dot_0MinusRc2SchemaDot represents a AsyncApi_2Dot_0Dot_0MinusRc2SchemaDot model.
type AsyncApi_2Dot_0Dot_0MinusRc2SchemaDot struct {
  Asyncapi *Asyncapi
  Id string
  Info *Info
  Servers map[string]Server
  DefaultContentType string
  Channels map[string]ChannelItem
  Components *Components
  Tags []Tag
  ExternalDocs *ExternalDocs
  AdditionalProperties map[string]interface{}
}