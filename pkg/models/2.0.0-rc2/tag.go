
package models

// Tag represents a Tag model.
type Tag struct {
  Name string `json:"name,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}