
package models
import (  
  "encoding/json"
)
// UserPasswordType represents an enum of UserPasswordType.
type UserPasswordType uint

const (
  UserPasswordTypeUserPassword UserPasswordType = iota
)

// Value returns the value of the enum.
func (op UserPasswordType) Value() any {
	if op >= UserPasswordType(len(UserPasswordTypeValues)) {
		return nil
	}
	return UserPasswordTypeValues[op]
}

var UserPasswordTypeValues = []any{"userPassword"}
var ValuesToUserPasswordType = map[any]UserPasswordType{
  UserPasswordTypeValues[UserPasswordTypeUserPassword]: UserPasswordTypeUserPassword,
}

 
func (op *UserPasswordType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToUserPasswordType[v]
  return nil
}

func (op UserPasswordType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}