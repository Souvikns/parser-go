
package models
import (  
  "encoding/json"
)
// BearerHttpSecuritySchemeScheme represents an enum of BearerHttpSecuritySchemeScheme.
type BearerHttpSecuritySchemeScheme uint

const (
  BearerHttpSecuritySchemeSchemeBearer BearerHttpSecuritySchemeScheme = iota
)

// Value returns the value of the enum.
func (op BearerHttpSecuritySchemeScheme) Value() any {
	if op >= BearerHttpSecuritySchemeScheme(len(BearerHttpSecuritySchemeSchemeValues)) {
		return nil
	}
	return BearerHttpSecuritySchemeSchemeValues[op]
}

var BearerHttpSecuritySchemeSchemeValues = []any{"bearer"}
var ValuesToBearerHttpSecuritySchemeScheme = map[any]BearerHttpSecuritySchemeScheme{
  BearerHttpSecuritySchemeSchemeValues[BearerHttpSecuritySchemeSchemeBearer]: BearerHttpSecuritySchemeSchemeBearer,
}

 
          
func (op *BearerHttpSecuritySchemeScheme) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBearerHttpSecuritySchemeScheme[v]
	return nil
}

func (op BearerHttpSecuritySchemeScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          