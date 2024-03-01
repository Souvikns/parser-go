
package models
import (  
  "encoding/json"
)
// MessageBindingsObjectJmsBindingVersion represents an enum of MessageBindingsObjectJmsBindingVersion.
type MessageBindingsObjectJmsBindingVersion uint

const (
  MessageBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1 MessageBindingsObjectJmsBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectJmsBindingVersion) Value() any {
	if op >= MessageBindingsObjectJmsBindingVersion(len(MessageBindingsObjectJmsBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectJmsBindingVersionValues[op]
}

var MessageBindingsObjectJmsBindingVersionValues = []any{"0.0.1"}
var ValuesToMessageBindingsObjectJmsBindingVersion = map[any]MessageBindingsObjectJmsBindingVersion{
  MessageBindingsObjectJmsBindingVersionValues[MessageBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1]: MessageBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1,
}

 
          
func (op *MessageBindingsObjectJmsBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToMessageBindingsObjectJmsBindingVersion[v]
	return nil
}

func (op MessageBindingsObjectJmsBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          