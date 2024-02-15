
package models

// Channel represents a Channel model.
type Channel struct {
  Address string
  Messages map[string]interface{}
  Parameters map[string]interface{}
  Title string
  Summary string
  Description string
  Servers []Reference
  Tags []interface{}
  ExternalDocs interface{}
  Bindings interface{}
  AdditionalProperties map[string]interface{}
}