
package models

// Info represents a Info model.
type Info struct {
  Title string `json:"title"`
  Version string `json:"version"`
  Description string `json:"description"`
  TermsOfService string `json:"termsOfService"`
  Contact *Contact `json:"contact"`
  License *License `json:"license"`
  Tags []InfoTagsItem `json:"tags"`
  ExternalDocs *InfoExternalDocs `json:"externalDocs"`
  AdditionalProperties map[string]interface{} `json:"-"`
}