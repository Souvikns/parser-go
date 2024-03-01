
package models
import (  
  "encoding/json"
)
// BearerHttpSecuritySchemeType represents an enum of BearerHttpSecuritySchemeType.
type BearerHttpSecuritySchemeType uint

const (
  BearerHttpSecuritySchemeTypeHttp BearerHttpSecuritySchemeType = iota
)

// Value returns the value of the enum.
func (op BearerHttpSecuritySchemeType) Value() any {
	if op >= BearerHttpSecuritySchemeType(len(BearerHttpSecuritySchemeTypeValues)) {
		return nil
	}
	return BearerHttpSecuritySchemeTypeValues[op]
}

var BearerHttpSecuritySchemeTypeValues = []any{"http"}
var ValuesToBearerHttpSecuritySchemeType = map[any]BearerHttpSecuritySchemeType{
  BearerHttpSecuritySchemeTypeValues[BearerHttpSecuritySchemeTypeHttp]: BearerHttpSecuritySchemeTypeHttp,
}

 
          
func (op *BearerHttpSecuritySchemeType) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBearerHttpSecuritySchemeType[v]
	return nil
}

func (op BearerHttpSecuritySchemeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          