
package models

// Info represents a Info model.
type Info struct {
  Title string `json:"title,omitempty"`
  Version string `json:"version,omitempty"`
  Description string `json:"description,omitempty"`
  TermsOfService string `json:"termsOfService,omitempty"`
  Contact *Contact `json:"contact,omitempty"`
  License *License `json:"license,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}