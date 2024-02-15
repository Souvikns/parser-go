
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  SchemaFormat string
  ContentType string
  Headers interface{}
  MessageId string
  CorrelationId interface{}
  Tags []Tag
  Summary string
  Name string
  Title string
  Description string
  ExternalDocs *ExternalDocs
  Deprecated bool
  Examples []map[string]interface{}
  Bindings *BindingsObject
  AdditionalProperties map[string]interface{}
}