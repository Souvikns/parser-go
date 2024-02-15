
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  ContentType string
  Headers interface{}
  CorrelationId interface{}
  Tags []interface{}
  Summary string
  Name string
  Title string
  Description string
  ExternalDocs interface{}
  Deprecated bool
  Examples []map[string]interface{}
  Bindings interface{}
  AdditionalProperties map[string]interface{}
}