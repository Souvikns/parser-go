
package models

// SymmetricEncryption represents a SymmetricEncryption model.
type SymmetricEncryption struct {
  ReservedType *SymmetricEncryptionType `json:"type"`
  Description string `json:"description"`
  AdditionalProperties map[string]interface{} `json:"-"`
}