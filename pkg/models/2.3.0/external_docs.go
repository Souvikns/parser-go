
package models

// ExternalDocs represents a ExternalDocs model.
type ExternalDocs struct {
  Description string `json:"description,omitempty"`
  Url string `json:"url,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-",omitempty`
}