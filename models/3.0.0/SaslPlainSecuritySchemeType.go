
package models
import (  
  "encoding/json"
)
// SaslPlainSecuritySchemeType represents an enum of SaslPlainSecuritySchemeType.
type SaslPlainSecuritySchemeType uint

const (
  SaslPlainSecuritySchemeTypePlain SaslPlainSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op SaslPlainSecuritySchemeType) Value() any {
	if op >= SaslPlainSecuritySchemeType(len(SaslPlainSecuritySchemeTypeValues)) {
		return nil
	}
	return SaslPlainSecuritySchemeTypeValues[op]
}

var SaslPlainSecuritySchemeTypeValues = []any{"plain"}
var ValuesToSaslPlainSecuritySchemeType = map[any]SaslPlainSecuritySchemeType{
  SaslPlainSecuritySchemeTypeValues[SaslPlainSecuritySchemeTypePlain]: SaslPlainSecuritySchemeTypePlain,
}

 
          
func (op *SaslPlainSecuritySchemeType) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToSaslPlainSecuritySchemeType[v]
	return nil
}

func (op SaslPlainSecuritySchemeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          