
package models

// Tag represents a Tag model.
type Tag struct {
  Name string `json:"name"`
  Description string `json:"description"`
  ExternalDocs interface{} `json:"externalDocs"`
  AdditionalProperties map[string]interface{} `json:"-"`
}