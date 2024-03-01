
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectGooglepubsubBindingVersion represents an enum of ChannelBindingsObjectGooglepubsubBindingVersion.
type ChannelBindingsObjectGooglepubsubBindingVersion uint

const (
  ChannelBindingsObjectGooglepubsubBindingVersionNumber_0Dot_2Dot_0 ChannelBindingsObjectGooglepubsubBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectGooglepubsubBindingVersion) Value() any {
	if op >= ChannelBindingsObjectGooglepubsubBindingVersion(len(ChannelBindingsObjectGooglepubsubBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectGooglepubsubBindingVersionValues[op]
}

var ChannelBindingsObjectGooglepubsubBindingVersionValues = []any{"0.2.0"}
var ValuesToChannelBindingsObjectGooglepubsubBindingVersion = map[any]ChannelBindingsObjectGooglepubsubBindingVersion{
  ChannelBindingsObjectGooglepubsubBindingVersionValues[ChannelBindingsObjectGooglepubsubBindingVersionNumber_0Dot_2Dot_0]: ChannelBindingsObjectGooglepubsubBindingVersionNumber_0Dot_2Dot_0,
}

 
          
func (op *ChannelBindingsObjectGooglepubsubBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToChannelBindingsObjectGooglepubsubBindingVersion[v]
	return nil
}

func (op ChannelBindingsObjectGooglepubsubBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          