
package models

// MessageBindingsObjectAnypointmqBindingVersion represents an enum of MessageBindingsObjectAnypointmqBindingVersion.
type MessageBindingsObjectAnypointmqBindingVersion uint

const (
  MessageBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1 MessageBindingsObjectAnypointmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectAnypointmqBindingVersion) Value() any {
	if op >= MessageBindingsObjectAnypointmqBindingVersion(len(MessageBindingsObjectAnypointmqBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectAnypointmqBindingVersionValues[op]
}

var MessageBindingsObjectAnypointmqBindingVersionValues = []any{"0.0.1"}
var ValuesToMessageBindingsObjectAnypointmqBindingVersion = map[any]MessageBindingsObjectAnypointmqBindingVersion{
  MessageBindingsObjectAnypointmqBindingVersionValues[MessageBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1]: MessageBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1,
}
