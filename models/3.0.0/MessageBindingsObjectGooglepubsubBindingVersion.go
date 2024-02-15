
package models

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
