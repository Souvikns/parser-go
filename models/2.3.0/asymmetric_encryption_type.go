
package models
import (  
  "encoding/json"
)
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

 
func (op *AsymmetricEncryptionType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToAsymmetricEncryptionType[v]
  return nil
}

func (op AsymmetricEncryptionType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}