
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectSqsBindingVersion represents an enum of ChannelBindingsObjectSqsBindingVersion.
type ChannelBindingsObjectSqsBindingVersion uint

const (
  ChannelBindingsObjectSqsBindingVersionNumber_0Dot_2Dot_0 ChannelBindingsObjectSqsBindingVersion = iota
  ChannelBindingsObjectSqsBindingVersionNumber_0Dot_1Dot_0
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectSqsBindingVersion) Value() any {
	if op >= ChannelBindingsObjectSqsBindingVersion(len(ChannelBindingsObjectSqsBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectSqsBindingVersionValues[op]
}

var ChannelBindingsObjectSqsBindingVersionValues = []any{"0.2.0","0.1.0"}
var ValuesToChannelBindingsObjectSqsBindingVersion = map[any]ChannelBindingsObjectSqsBindingVersion{
  ChannelBindingsObjectSqsBindingVersionValues[ChannelBindingsObjectSqsBindingVersionNumber_0Dot_2Dot_0]: ChannelBindingsObjectSqsBindingVersionNumber_0Dot_2Dot_0,
  ChannelBindingsObjectSqsBindingVersionValues[ChannelBindingsObjectSqsBindingVersionNumber_0Dot_1Dot_0]: ChannelBindingsObjectSqsBindingVersionNumber_0Dot_1Dot_0,
}

 
func (op *ChannelBindingsObjectSqsBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectSqsBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectSqsBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}