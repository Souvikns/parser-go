
package models

// SymmetricEncryptionType represents an enum of SymmetricEncryptionType.
type SymmetricEncryptionType uint

const (
  SymmetricEncryptionTypeSymmetricEncryption SymmetricEncryptionType = iota
)

// Value returns the value of the enum.
func (op SymmetricEncryptionType) Value() any {
	if op >= SymmetricEncryptionType(len(SymmetricEncryptionTypeValues)) {
		return nil
	}
	return SymmetricEncryptionTypeValues[op]
}

var SymmetricEncryptionTypeValues = []any{"symmetricEncryption"}
var ValuesToSymmetricEncryptionType = map[any]SymmetricEncryptionType{
  SymmetricEncryptionTypeValues[SymmetricEncryptionTypeSymmetricEncryption]: SymmetricEncryptionTypeSymmetricEncryption,
}
