
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectIbmmqBindingVersion represents an enum of ChannelBindingsObjectIbmmqBindingVersion.
type ChannelBindingsObjectIbmmqBindingVersion uint

const (
  ChannelBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0 ChannelBindingsObjectIbmmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectIbmmqBindingVersion) Value() any {
	if op >= ChannelBindingsObjectIbmmqBindingVersion(len(ChannelBindingsObjectIbmmqBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectIbmmqBindingVersionValues[op]
}

var ChannelBindingsObjectIbmmqBindingVersionValues = []any{"0.1.0"}
var ValuesToChannelBindingsObjectIbmmqBindingVersion = map[any]ChannelBindingsObjectIbmmqBindingVersion{
  ChannelBindingsObjectIbmmqBindingVersionValues[ChannelBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0]: ChannelBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0,
}

 
func (op *ChannelBindingsObjectIbmmqBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectIbmmqBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectIbmmqBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}