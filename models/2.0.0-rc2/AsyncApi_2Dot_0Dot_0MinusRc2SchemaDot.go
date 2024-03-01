
package models

// AsyncApi_2Dot_0Dot_0MinusRc2SchemaDot represents a AsyncApi_2Dot_0Dot_0MinusRc2SchemaDot model.
type AsyncApi_2Dot_0Dot_0MinusRc2SchemaDot struct {
  Asyncapi *Asyncapi `json:asyncapi`
  Id string `json:id`
  Info *Info `json:info`
  Servers map[string]Server `json:servers`
  DefaultContentType string `json:defaultContentType`
  Channels map[string]ChannelItem `json:channels`
  Components *Components `json:components`
  Tags []Tag `json:tags`
  ExternalDocs *ExternalDocs `json:externalDocs`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}