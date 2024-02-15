
package models

// ServerBindingsObjectIbmmqBindingVersion represents an enum of ServerBindingsObjectIbmmqBindingVersion.
type ServerBindingsObjectIbmmqBindingVersion uint

const (
  ServerBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0 ServerBindingsObjectIbmmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op ServerBindingsObjectIbmmqBindingVersion) Value() any {
	if op >= ServerBindingsObjectIbmmqBindingVersion(len(ServerBindingsObjectIbmmqBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectIbmmqBindingVersionValues[op]
}

var ServerBindingsObjectIbmmqBindingVersionValues = []any{"0.1.0"}
var ValuesToServerBindingsObjectIbmmqBindingVersion = map[any]ServerBindingsObjectIbmmqBindingVersion{
  ServerBindingsObjectIbmmqBindingVersionValues[ServerBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0]: ServerBindingsObjectIbmmqBindingVersionNumber_0Dot_1Dot_0,
}
