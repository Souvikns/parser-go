
package models

// MessageBindingsObjectMqttBindingVersion represents an enum of MessageBindingsObjectMqttBindingVersion.
type MessageBindingsObjectMqttBindingVersion uint

const (
  MessageBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0 MessageBindingsObjectMqttBindingVersion = iota
)

// Value returns the value of the enum.
func (op MessageBindingsObjectMqttBindingVersion) Value() any {
	if op >= MessageBindingsObjectMqttBindingVersion(len(MessageBindingsObjectMqttBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectMqttBindingVersionValues[op]
}

var MessageBindingsObjectMqttBindingVersionValues = []any{"0.2.0"}
var ValuesToMessageBindingsObjectMqttBindingVersion = map[any]MessageBindingsObjectMqttBindingVersion{
  MessageBindingsObjectMqttBindingVersionValues[MessageBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0]: MessageBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0,
}
