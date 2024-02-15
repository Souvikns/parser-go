
package models

// Message represents a Message model.
type Message struct {
  SchemaFormat string
  ContentType string
  Headers map[string]interface{}
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
  ProtocolInfo map[string]map[string]interface{}
  Traits []interface{}
  AdditionalProperties map[string]interface{}
}