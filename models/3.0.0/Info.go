
package models

// Info represents a Info model.
type Info struct {
  Title string
  Version string
  Description string
  TermsOfService string
  Contact *Contact
  License *License
  Tags []interface{}
  ExternalDocs interface{}
  AdditionalProperties map[string]interface{}
}