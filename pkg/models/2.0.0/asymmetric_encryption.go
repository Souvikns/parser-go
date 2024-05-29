
package models

// AsymmetricEncryption represents a AsymmetricEncryption model.
type AsymmetricEncryption struct {
  ReservedType *AsymmetricEncryptionType `json:"type,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}