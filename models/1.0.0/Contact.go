
package models

// Contact represents a Contact model.
type Contact struct {
  Name string `json:name`
  Url string `json:url`
  Email string `json:email`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}