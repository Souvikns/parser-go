
package models
import (  
  "encoding/json"
)
// X509Type represents an enum of X509Type.
type X509Type uint

const (
  X509TypeX509 X509Type = iota
)

// Value returns the value of the enum.
func (op X509Type) Value() any {
	if op >= X509Type(len(X509TypeValues)) {
		return nil
	}
	return X509TypeValues[op]
}

var X509TypeValues = []any{"X509"}
var ValuesToX509Type = map[any]X509Type{
  X509TypeValues[X509TypeX509]: X509TypeX509,
}

 
          
func (op *X509Type) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToX509Type[v]
	return nil
}

func (op X509Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          