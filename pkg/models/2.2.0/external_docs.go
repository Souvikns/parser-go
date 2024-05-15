
package models

// ExternalDocs represents a ExternalDocs model.
type ExternalDocs struct {
  Description string `json:"description"`
  Url string `json:"url"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}