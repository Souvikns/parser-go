
package models

// ServerBindingsObjectPulsarBindingVersion represents an enum of ServerBindingsObjectPulsarBindingVersion.
type ServerBindingsObjectPulsarBindingVersion uint

const (
  ServerBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0 ServerBindingsObjectPulsarBindingVersion = iota
)

// Value returns the value of the enum.
func (op ServerBindingsObjectPulsarBindingVersion) Value() any {
	if op >= ServerBindingsObjectPulsarBindingVersion(len(ServerBindingsObjectPulsarBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectPulsarBindingVersionValues[op]
}

var ServerBindingsObjectPulsarBindingVersionValues = []any{"0.1.0"}
var ValuesToServerBindingsObjectPulsarBindingVersion = map[any]ServerBindingsObjectPulsarBindingVersion{
  ServerBindingsObjectPulsarBindingVersionValues[ServerBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0]: ServerBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0,
}
