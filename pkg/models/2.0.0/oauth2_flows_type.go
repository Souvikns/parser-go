
package models
import (  
  "encoding/json"
)
// Oauth2FlowsType represents an enum of Oauth2FlowsType.
type Oauth2FlowsType uint

const (
  Oauth2FlowsTypeOauth2 Oauth2FlowsType = iota
)

// Value returns the value of the enum.
func (op Oauth2FlowsType) Value() any {
	if op >= Oauth2FlowsType(len(Oauth2FlowsTypeValues)) {
		return nil
	}
	return Oauth2FlowsTypeValues[op]
}

var Oauth2FlowsTypeValues = []any{"oauth2"}
var ValuesToOauth2FlowsType = map[any]Oauth2FlowsType{
  Oauth2FlowsTypeValues[Oauth2FlowsTypeOauth2]: Oauth2FlowsTypeOauth2,
}

 
func (op *Oauth2FlowsType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToOauth2FlowsType[v]
  return nil
}

func (op Oauth2FlowsType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}