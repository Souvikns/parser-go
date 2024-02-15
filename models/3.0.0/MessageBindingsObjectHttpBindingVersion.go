
package models

// MessageBindingsObjectHttpBindingVersion represents an enum of MessageBindingsObjectHttpBindingVersion.
type MessageBindingsObjectHttpBindingVersion uint

const (
  MessageBindingsObjectHttpBindingVersionNumber_0Dot_2Dot_0 MessageBindingsObjectHttpBindingVersion = iota
  MessageBindingsObjectHttpBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op MessageBindingsObjectHttpBindingVersion) Value() any {
	if op >= MessageBindingsObjectHttpBindingVersion(len(MessageBindingsObjectHttpBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectHttpBindingVersionValues[op]
}

var MessageBindingsObjectHttpBindingVersionValues = []any{"0.2.0","0.3.0"}
var ValuesToMessageBindingsObjectHttpBindingVersion = map[any]MessageBindingsObjectHttpBindingVersion{
  MessageBindingsObjectHttpBindingVersionValues[MessageBindingsObjectHttpBindingVersionNumber_0Dot_2Dot_0]: MessageBindingsObjectHttpBindingVersionNumber_0Dot_2Dot_0,
  MessageBindingsObjectHttpBindingVersionValues[MessageBindingsObjectHttpBindingVersionNumber_0Dot_3Dot_0]: MessageBindingsObjectHttpBindingVersionNumber_0Dot_3Dot_0,
}
