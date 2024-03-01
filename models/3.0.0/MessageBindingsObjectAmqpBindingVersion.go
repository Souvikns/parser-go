
package models
import (  
  "encoding/json"
)
// MessageBindingsObjectAmqpBindingVersion represents an enum of MessageBindingsObjectAmqpBindingVersion.
type MessageBindingsObjectAmqpBindingVersion uint

const (
  MessageBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0 MessageBindingsObjectAmqpBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectAmqpBindingVersion) Value() any {
	if op >= MessageBindingsObjectAmqpBindingVersion(len(MessageBindingsObjectAmqpBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectAmqpBindingVersionValues[op]
}

var MessageBindingsObjectAmqpBindingVersionValues = []any{"0.3.0"}
var ValuesToMessageBindingsObjectAmqpBindingVersion = map[any]MessageBindingsObjectAmqpBindingVersion{
  MessageBindingsObjectAmqpBindingVersionValues[MessageBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0]: MessageBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0,
}

 
          
func (op *MessageBindingsObjectAmqpBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToMessageBindingsObjectAmqpBindingVersion[v]
	return nil
}

func (op MessageBindingsObjectAmqpBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          