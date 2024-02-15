
package models

// MessageBindingsObjectKafkaBindingVersion represents an enum of MessageBindingsObjectKafkaBindingVersion.
type MessageBindingsObjectKafkaBindingVersion uint

const (
  MessageBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0 MessageBindingsObjectKafkaBindingVersion = iota
  MessageBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op MessageBindingsObjectKafkaBindingVersion) Value() any {
	if op >= MessageBindingsObjectKafkaBindingVersion(len(MessageBindingsObjectKafkaBindingVersionValues)) {
		return nil
	}
	return MessageBindingsObjectKafkaBindingVersionValues[op]
}

var MessageBindingsObjectKafkaBindingVersionValues = []any{"0.4.0","0.3.0"}
var ValuesToMessageBindingsObjectKafkaBindingVersion = map[any]MessageBindingsObjectKafkaBindingVersion{
  MessageBindingsObjectKafkaBindingVersionValues[MessageBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0]: MessageBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0,
  MessageBindingsObjectKafkaBindingVersionValues[MessageBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0]: MessageBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0,
}
