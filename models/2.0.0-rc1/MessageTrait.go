
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  SchemaFormat string
  ContentType string
  Headers map[string]interface{}
  CorrelationId interface{}
  Tags []Tag
  Summary string
  Name string
  Title string
  Description string
  ExternalDocs *ExternalDocs
  Deprecated bool
  Examples []map[string]interface{}
  ProtocolInfo map[string]map[string]interface{}
  AdditionalProperties map[string]interface{}
}