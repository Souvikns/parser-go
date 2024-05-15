
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectAmqpBindingVersion represents an enum of ChannelBindingsObjectAmqpBindingVersion.
type ChannelBindingsObjectAmqpBindingVersion uint

const (
  ChannelBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0 ChannelBindingsObjectAmqpBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectAmqpBindingVersion) Value() any {
	if op >= ChannelBindingsObjectAmqpBindingVersion(len(ChannelBindingsObjectAmqpBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectAmqpBindingVersionValues[op]
}

var ChannelBindingsObjectAmqpBindingVersionValues = []any{"0.3.0"}
var ValuesToChannelBindingsObjectAmqpBindingVersion = map[any]ChannelBindingsObjectAmqpBindingVersion{
  ChannelBindingsObjectAmqpBindingVersionValues[ChannelBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0]: ChannelBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0,
}

 
func (op *ChannelBindingsObjectAmqpBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectAmqpBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectAmqpBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}