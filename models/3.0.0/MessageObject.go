
package models

// MessageObject represents a MessageObject model.
type MessageObject struct {
  ContentType string
  Headers interface{}
  Payload interface{}
  CorrelationId interface{}
  Tags []interface{}
  Summary string
  Name string
  Title string
  Description string
  ExternalDocs interface{}
  Deprecated bool
  Examples []interface{}
  Bindings interface{}
  Traits []interface{}
  AdditionalProperties map[string]interface{}
}