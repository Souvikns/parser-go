
package models

// License represents a License model.
type License struct {
  Name string `json:"name,omitempty"`
  Url string `json:"url,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}