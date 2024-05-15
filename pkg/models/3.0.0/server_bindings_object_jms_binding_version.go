
package models
import (  
  "encoding/json"
)
// ServerBindingsObjectJmsBindingVersion represents an enum of ServerBindingsObjectJmsBindingVersion.
type ServerBindingsObjectJmsBindingVersion uint

const (
  ServerBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1 ServerBindingsObjectJmsBindingVersion = iota
)

// Value returns the value of the enum.
func (op ServerBindingsObjectJmsBindingVersion) Value() any {
	if op >= ServerBindingsObjectJmsBindingVersion(len(ServerBindingsObjectJmsBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectJmsBindingVersionValues[op]
}

var ServerBindingsObjectJmsBindingVersionValues = []any{"0.0.1"}
var ValuesToServerBindingsObjectJmsBindingVersion = map[any]ServerBindingsObjectJmsBindingVersion{
  ServerBindingsObjectJmsBindingVersionValues[ServerBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1]: ServerBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1,
}

 
func (op *ServerBindingsObjectJmsBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToServerBindingsObjectJmsBindingVersion[v]
  return nil
}

func (op ServerBindingsObjectJmsBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}