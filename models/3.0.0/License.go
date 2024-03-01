
package models

// License represents a License model.
type License struct {
  Name string `json:"name"`
  Url string `json:"url"`
  AdditionalProperties map[string]interface{} `json:"-"`
}