
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectSnsBindingVersion represents an enum of ChannelBindingsObjectSnsBindingVersion.
type ChannelBindingsObjectSnsBindingVersion uint

const (
  ChannelBindingsObjectSnsBindingVersionNumber_0Dot_1Dot_0 ChannelBindingsObjectSnsBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectSnsBindingVersion) Value() any {
	if op >= ChannelBindingsObjectSnsBindingVersion(len(ChannelBindingsObjectSnsBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectSnsBindingVersionValues[op]
}

var ChannelBindingsObjectSnsBindingVersionValues = []any{"0.1.0"}
var ValuesToChannelBindingsObjectSnsBindingVersion = map[any]ChannelBindingsObjectSnsBindingVersion{
  ChannelBindingsObjectSnsBindingVersionValues[ChannelBindingsObjectSnsBindingVersionNumber_0Dot_1Dot_0]: ChannelBindingsObjectSnsBindingVersionNumber_0Dot_1Dot_0,
}

 
func (op *ChannelBindingsObjectSnsBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectSnsBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectSnsBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}