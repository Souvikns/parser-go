
package models

// MessageOneOf_1OneOf_1 represents a MessageOneOf_1OneOf_1 model.
type MessageOneOf_1OneOf_1 struct {
  SchemaFormat string
  ContentType string
  Headers interface{}
  Payload interface{}
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
  Traits []interface{}
  AdditionalProperties map[string]interface{}
}