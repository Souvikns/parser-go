
package models
import (  
  "encoding/json"
)
// MessageBindingsObjectAnypointmqBindingVersion represents an enum of MessageBindingsObjectAnypointmqBindingVersion.
type MessageBindingsObjectAnypointmqBindingVersion uint

const (
  MessageBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1 MessageBindingsObjectAnypointmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectAnypointmqBindingVersion) Value() any {
	if op >= MessageBindingsObjectAnypointmqBindingVersion(len(MessageBindingsObjectAnypointmqBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectAnypointmqBindingVersionValues[op]
}

var MessageBindingsObjectAnypointmqBindingVersionValues = []any{"0.0.1"}
var ValuesToMessageBindingsObjectAnypointmqBindingVersion = map[any]MessageBindingsObjectAnypointmqBindingVersion{
  MessageBindingsObjectAnypointmqBindingVersionValues[MessageBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1]: MessageBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1,
}

 
func (op *MessageBindingsObjectAnypointmqBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToMessageBindingsObjectAnypointmqBindingVersion[v]
  return nil
}

func (op MessageBindingsObjectAnypointmqBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}