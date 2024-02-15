
package models

// ServerBindingsObjectMqttBindingVersion represents an enum of ServerBindingsObjectMqttBindingVersion.
type ServerBindingsObjectMqttBindingVersion uint

const (
  ServerBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0 ServerBindingsObjectMqttBindingVersion = iota
)

// Value returns the value of the enum.
func (op ServerBindingsObjectMqttBindingVersion) Value() any {
	if op >= ServerBindingsObjectMqttBindingVersion(len(ServerBindingsObjectMqttBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectMqttBindingVersionValues[op]
}

var ServerBindingsObjectMqttBindingVersionValues = []any{"0.2.0"}
var ValuesToServerBindingsObjectMqttBindingVersion = map[any]ServerBindingsObjectMqttBindingVersion{
  ServerBindingsObjectMqttBindingVersionValues[ServerBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0]: ServerBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0,
}
