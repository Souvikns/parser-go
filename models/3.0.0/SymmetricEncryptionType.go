
package models
import (  
  "encoding/json"
)
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

 
          
func (op *SymmetricEncryptionType) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToSymmetricEncryptionType[v]
	return nil
}

func (op SymmetricEncryptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          