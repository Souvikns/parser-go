
package models
import (  
  "encoding/json"
)
// MessageBindingsObjectIbmmqBindingVersion represents an enum of MessageBindingsObjectIbmmqBindingVersion.
type MessageBindingsObjectIbmmqBindingVersion uint

const (
  MessageBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0 MessageBindingsObjectIbmmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectIbmmqBindingVersion) Value() any {
	if op >= MessageBindingsObjectIbmmqBindingVersion(len(MessageBindingsObjectIbmmqBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectIbmmqBindingVersionValues[op]
}

var MessageBindingsObjectIbmmqBindingVersionValues = []any{"0.1.0"}
var ValuesToMessageBindingsObjectIbmmqBindingVersion = map[any]MessageBindingsObjectIbmmqBindingVersion{
  MessageBindingsObjectIbmmqBindingVersionValues[MessageBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0]: MessageBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0,
}

 
          
func (op *MessageBindingsObjectIbmmqBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToMessageBindingsObjectIbmmqBindingVersion[v]
	return nil
}

func (op MessageBindingsObjectIbmmqBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          