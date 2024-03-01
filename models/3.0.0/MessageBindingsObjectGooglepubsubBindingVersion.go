
package models
import (  
  "encoding/json"
)
// MessageBindingsObjectGooglepubsubBindingVersion represents an enum of MessageBindingsObjectGooglepubsubBindingVersion.
type MessageBindingsObjectGooglepubsubBindingVersion uint

const (
  MessageBindingsObjectGooglepubsubBindingVersionNumber_0Dot_2Dot_0 MessageBindingsObjectGooglepubsubBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectGooglepubsubBindingVersion) Value() any {
	if op >= MessageBindingsObjectGooglepubsubBindingVersion(len(MessageBindingsObjectGooglepubsubBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectGooglepubsubBindingVersionValues[op]
}

var MessageBindingsObjectGooglepubsubBindingVersionValues = []any{"0.2.0"}
var ValuesToMessageBindingsObjectGooglepubsubBindingVersion = map[any]MessageBindingsObjectGooglepubsubBindingVersion{
  MessageBindingsObjectGooglepubsubBindingVersionValues[MessageBindingsObjectGooglepubsubBindingVersionNumber_0Dot_2Dot_0]: MessageBindingsObjectGooglepubsubBindingVersionNumber_0Dot_2Dot_0,
}

 
          
func (op *MessageBindingsObjectGooglepubsubBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToMessageBindingsObjectGooglepubsubBindingVersion[v]
	return nil
}

func (op MessageBindingsObjectGooglepubsubBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          