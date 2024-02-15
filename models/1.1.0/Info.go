
package models

// Info represents a Info model.
type Info struct {
  Title string
  Version string
  Description string
  TermsOfService string
  Contact *Contact
  License *License
  AdditionalProperties map[string]interface{}
}