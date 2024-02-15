
package models

// OperationBindingsObjectMqttBindingVersion represents an enum of OperationBindingsObjectMqttBindingVersion.
type OperationBindingsObjectMqttBindingVersion uint

const (
  OperationBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0 OperationBindingsObjectMqttBindingVersion = iota
)

// Value returns the value of the enum.
func (op OperationBindingsObjectMqttBindingVersion) Value() any {
	if op >= OperationBindingsObjectMqttBindingVersion(len(OperationBindingsObjectMqttBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectMqttBindingVersionValues[op]
}

var OperationBindingsObjectMqttBindingVersionValues = []any{"0.2.0"}
var ValuesToOperationBindingsObjectMqttBindingVersion = map[any]OperationBindingsObjectMqttBindingVersion{
  OperationBindingsObjectMqttBindingVersionValues[OperationBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0]: OperationBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0,
}
