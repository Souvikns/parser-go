
package models

// Contact represents a Contact model.
type Contact struct {
  Name string `json:"name,omitempty"`
  Url string `json:"url,omitempty"`
  Email string `json:"email,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}