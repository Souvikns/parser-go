
package models

// X509 represents a X509 model.
type X509 struct {
  ReservedType *X509Type `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}