
package models

// AsymmetricEncryption represents a AsymmetricEncryption model.
type AsymmetricEncryption struct {
  ReservedType *AsymmetricEncryptionType `json:type`
  Description string `json:description`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}