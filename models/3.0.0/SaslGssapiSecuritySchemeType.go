
package models
import (  
  "encoding/json"
)
// SaslGssapiSecuritySchemeType represents an enum of SaslGssapiSecuritySchemeType.
type SaslGssapiSecuritySchemeType uint

const (
  SaslGssapiSecuritySchemeTypeGssapi SaslGssapiSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op SaslGssapiSecuritySchemeType) Value() any {
	if op >= SaslGssapiSecuritySchemeType(len(SaslGssapiSecuritySchemeTypeValues)) {
		return nil
	}
	return SaslGssapiSecuritySchemeTypeValues[op]
}

var SaslGssapiSecuritySchemeTypeValues = []any{"gssapi"}
var ValuesToSaslGssapiSecuritySchemeType = map[any]SaslGssapiSecuritySchemeType{
  SaslGssapiSecuritySchemeTypeValues[SaslGssapiSecuritySchemeTypeGssapi]: SaslGssapiSecuritySchemeTypeGssapi,
}

 
          
func (op *SaslGssapiSecuritySchemeType) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToSaslGssapiSecuritySchemeType[v]
	return nil
}

func (op SaslGssapiSecuritySchemeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          