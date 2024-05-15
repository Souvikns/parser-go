
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectPulsarBindingVersion represents an enum of ChannelBindingsObjectPulsarBindingVersion.
type ChannelBindingsObjectPulsarBindingVersion uint

const (
  ChannelBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0 ChannelBindingsObjectPulsarBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectPulsarBindingVersion) Value() any {
	if op >= ChannelBindingsObjectPulsarBindingVersion(len(ChannelBindingsObjectPulsarBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectPulsarBindingVersionValues[op]
}

var ChannelBindingsObjectPulsarBindingVersionValues = []any{"0.1.0"}
var ValuesToChannelBindingsObjectPulsarBindingVersion = map[any]ChannelBindingsObjectPulsarBindingVersion{
  ChannelBindingsObjectPulsarBindingVersionValues[ChannelBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0]: ChannelBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0,
}

 
func (op *ChannelBindingsObjectPulsarBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectPulsarBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectPulsarBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}