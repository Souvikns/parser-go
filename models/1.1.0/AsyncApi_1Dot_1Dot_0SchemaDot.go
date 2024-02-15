
package models

// AsyncApi_1Dot_1Dot_0SchemaDot represents a AsyncApi_1Dot_1Dot_0SchemaDot model.
type AsyncApi_1Dot_1Dot_0SchemaDot struct {
  Asyncapi *Asyncapi
  Info *Info
  BaseTopic string
  Servers []Server
  Topics *Topics
  Components *Components
  Tags []Tag
  Security []map[string][]string
  ExternalDocs *ExternalDocs
  AdditionalProperties map[string]interface{}
}