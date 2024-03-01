
package models

// AsyncApi_1Dot_0SchemaDot represents a AsyncApi_1Dot_0SchemaDot model.
type AsyncApi_1Dot_0SchemaDot struct {
  Asyncapi *Asyncapi `json:asyncapi`
  Info *Info `json:info`
  BaseTopic string `json:baseTopic`
  Servers []Server `json:servers`
  Topics *Topics `json:topics`
  Components *Components `json:components`
  Tags []Tag `json:tags`
  Security []map[string][]string `json:security`
  ExternalDocs *ExternalDocs `json:externalDocs`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}