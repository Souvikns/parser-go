
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectAnypointmqBindingVersion represents an enum of ChannelBindingsObjectAnypointmqBindingVersion.
type ChannelBindingsObjectAnypointmqBindingVersion uint

const (
  ChannelBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1 ChannelBindingsObjectAnypointmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectAnypointmqBindingVersion) Value() any {
	if op >= ChannelBindingsObjectAnypointmqBindingVersion(len(ChannelBindingsObjectAnypointmqBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectAnypointmqBindingVersionValues[op]
}

var ChannelBindingsObjectAnypointmqBindingVersionValues = []any{"0.0.1"}
var ValuesToChannelBindingsObjectAnypointmqBindingVersion = map[any]ChannelBindingsObjectAnypointmqBindingVersion{
  ChannelBindingsObjectAnypointmqBindingVersionValues[ChannelBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1]: ChannelBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1,
}

 
func (op *ChannelBindingsObjectAnypointmqBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectAnypointmqBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectAnypointmqBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}