
package models

// AsymmetricEncryptionType represents an enum of AsymmetricEncryptionType.
type AsymmetricEncryptionType uint

const (
  AsymmetricEncryptionTypeAsymmetricEncryption AsymmetricEncryptionType = iota
)

// Value returns the value of the enum.
func (op AsymmetricEncryptionType) Value() any {
	if op >= AsymmetricEncryptionType(len(AsymmetricEncryptionTypeValues)) {
		return nil
	}
	return AsymmetricEncryptionTypeValues[op]
}

var AsymmetricEncryptionTypeValues = []any{"asymmetricEncryption"}
var ValuesToAsymmetricEncryptionType = map[any]AsymmetricEncryptionType{
  AsymmetricEncryptionTypeValues[AsymmetricEncryptionTypeAsymmetricEncryption]: AsymmetricEncryptionTypeAsymmetricEncryption,
}
