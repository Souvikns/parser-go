
package models
import (  
  "encoding/json"
)
// ChannelBindingsObjectKafkaBindingVersion represents an enum of ChannelBindingsObjectKafkaBindingVersion.
type ChannelBindingsObjectKafkaBindingVersion uint

const (
  ChannelBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0 ChannelBindingsObjectKafkaBindingVersion = iota
  ChannelBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectKafkaBindingVersion) Value() any {
	if op >= ChannelBindingsObjectKafkaBindingVersion(len(ChannelBindingsObjectKafkaBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectKafkaBindingVersionValues[op]
}

var ChannelBindingsObjectKafkaBindingVersionValues = []any{"0.4.0","0.3.0"}
var ValuesToChannelBindingsObjectKafkaBindingVersion = map[any]ChannelBindingsObjectKafkaBindingVersion{
  ChannelBindingsObjectKafkaBindingVersionValues[ChannelBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0]: ChannelBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0,
  ChannelBindingsObjectKafkaBindingVersionValues[ChannelBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0]: ChannelBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0,
}

 
func (op *ChannelBindingsObjectKafkaBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToChannelBindingsObjectKafkaBindingVersion[v]
  return nil
}

func (op ChannelBindingsObjectKafkaBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}