
package models

// Message represents a Message model.
type Message struct {
  Ref string
  Headers *Schema
  Payload *Schema
  Tags []Tag
  Summary string
  Description string
  ExternalDocs *ExternalDocs
  Deprecated bool
  Example interface{}
  AdditionalProperties map[string]interface{}
}