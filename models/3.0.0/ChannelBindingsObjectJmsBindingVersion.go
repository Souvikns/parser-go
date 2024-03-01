
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectJmsBindingVersion represents an enum of ChannelBindingsObjectJmsBindingVersion.
type ChannelBindingsObjectJmsBindingVersion uint

const (
  ChannelBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1 ChannelBindingsObjectJmsBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectJmsBindingVersion) Value() any {
	if op >= ChannelBindingsObjectJmsBindingVersion(len(ChannelBindingsObjectJmsBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectJmsBindingVersionValues[op]
}

var ChannelBindingsObjectJmsBindingVersionValues = []any{"0.0.1"}
var ValuesToChannelBindingsObjectJmsBindingVersion = map[any]ChannelBindingsObjectJmsBindingVersion{
  ChannelBindingsObjectJmsBindingVersionValues[ChannelBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1]: ChannelBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1,
}

 
          
func (op *ChannelBindingsObjectJmsBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToChannelBindingsObjectJmsBindingVersion[v]
	return nil
}

func (op ChannelBindingsObjectJmsBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          