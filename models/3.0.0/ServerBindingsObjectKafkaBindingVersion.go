
package models

// ServerBindingsObjectKafkaBindingVersion represents an enum of ServerBindingsObjectKafkaBindingVersion.
type ServerBindingsObjectKafkaBindingVersion uint

const (
  ServerBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0 ServerBindingsObjectKafkaBindingVersion = iota
  ServerBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op ServerBindingsObjectKafkaBindingVersion) Value() any {
	if op >= ServerBindingsObjectKafkaBindingVersion(len(ServerBindingsObjectKafkaBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectKafkaBindingVersionValues[op]
}

var ServerBindingsObjectKafkaBindingVersionValues = []any{"0.4.0","0.3.0"}
var ValuesToServerBindingsObjectKafkaBindingVersion = map[any]ServerBindingsObjectKafkaBindingVersion{
  ServerBindingsObjectKafkaBindingVersionValues[ServerBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0]: ServerBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0,
  ServerBindingsObjectKafkaBindingVersionValues[ServerBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0]: ServerBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0,
}
