
package models
import (  
  "encoding/json"
)
// ServerBindingsObjectMqttBindingVersion represents an enum of ServerBindingsObjectMqttBindingVersion.
type ServerBindingsObjectMqttBindingVersion uint

const (
  ServerBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0 ServerBindingsObjectMqttBindingVersion = iota
)

// Value returns the value of the enum.
func (op ServerBindingsObjectMqttBindingVersion) Value() any {
	if op >= ServerBindingsObjectMqttBindingVersion(len(ServerBindingsObjectMqttBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectMqttBindingVersionValues[op]
}

var ServerBindingsObjectMqttBindingVersionValues = []any{"0.2.0"}
var ValuesToServerBindingsObjectMqttBindingVersion = map[any]ServerBindingsObjectMqttBindingVersion{
  ServerBindingsObjectMqttBindingVersionValues[ServerBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0]: ServerBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0,
}

 
func (op *ServerBindingsObjectMqttBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToServerBindingsObjectMqttBindingVersion[v]
  return nil
}

func (op ServerBindingsObjectMqttBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}