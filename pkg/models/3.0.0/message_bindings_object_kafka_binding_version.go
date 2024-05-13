
package models
import (  
  "encoding/json"
)
// MessageBindingsObjectKafkaBindingVersion represents an enum of MessageBindingsObjectKafkaBindingVersion.
type MessageBindingsObjectKafkaBindingVersion uint

const (
  MessageBindingsObjectKafkaBindingVersionNumber_0Dot_5Dot_0 MessageBindingsObjectKafkaBindingVersion = iota
  MessageBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0
  MessageBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op MessageBindingsObjectKafkaBindingVersion) Value() any {
	if op >= MessageBindingsObjectKafkaBindingVersion(len(MessageBindingsObjectKafkaBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectKafkaBindingVersionValues[op]
}

var MessageBindingsObjectKafkaBindingVersionValues = []any{"0.5.0","0.4.0","0.3.0"}
var ValuesToMessageBindingsObjectKafkaBindingVersion = map[any]MessageBindingsObjectKafkaBindingVersion{
  MessageBindingsObjectKafkaBindingVersionValues[MessageBindingsObjectKafkaBindingVersionNumber_0Dot_5Dot_0]: MessageBindingsObjectKafkaBindingVersionNumber_0Dot_5Dot_0,
  MessageBindingsObjectKafkaBindingVersionValues[MessageBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0]: MessageBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0,
  MessageBindingsObjectKafkaBindingVersionValues[MessageBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0]: MessageBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0,
}

 
func (op *MessageBindingsObjectKafkaBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToMessageBindingsObjectKafkaBindingVersion[v]
  return nil
}

func (op MessageBindingsObjectKafkaBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}