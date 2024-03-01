
package models

// X509 represents a X509 model.
type X509 struct {
  ReservedType *X509Type `json:type`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}