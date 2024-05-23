
package models

// SymmetricEncryption represents a SymmetricEncryption model.
type SymmetricEncryption struct {
  ReservedType *SymmetricEncryptionType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}