
package models

// Tag represents a Tag model.
type Tag struct {
  Name string
  Description string
  ExternalDocs *ExternalDocs
  AdditionalProperties map[string]interface{}
}