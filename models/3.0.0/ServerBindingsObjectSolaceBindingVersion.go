
package models

// ServerBindingsObjectSolaceBindingVersion represents an enum of ServerBindingsObjectSolaceBindingVersion.
type ServerBindingsObjectSolaceBindingVersion uint

const (
  ServerBindingsObjectSolaceBindingVersionNumber_0Dot_4Dot_0 ServerBindingsObjectSolaceBindingVersion = iota
  ServerBindingsObjectSolaceBindingVersionNumber_0Dot_3Dot_0
  ServerBindingsObjectSolaceBindingVersionNumber_0Dot_2Dot_0
)

// Value returns the value of the enum.
func (op ServerBindingsObjectSolaceBindingVersion) Value() any {
	if op >= ServerBindingsObjectSolaceBindingVersion(len(ServerBindingsObjectSolaceBindingVersionValues)) {
		return nil
	}
	return ServerBindingsObjectSolaceBindingVersionValues[op]
}

var ServerBindingsObjectSolaceBindingVersionValues = []any{"0.4.0","0.3.0","0.2.0"}
var ValuesToServerBindingsObjectSolaceBindingVersion = map[any]ServerBindingsObjectSolaceBindingVersion{
  ServerBindingsObjectSolaceBindingVersionValues[ServerBindingsObjectSolaceBindingVersionNumber_0Dot_4Dot_0]: ServerBindingsObjectSolaceBindingVersionNumber_0Dot_4Dot_0,
  ServerBindingsObjectSolaceBindingVersionValues[ServerBindingsObjectSolaceBindingVersionNumber_0Dot_3Dot_0]: ServerBindingsObjectSolaceBindingVersionNumber_0Dot_3Dot_0,
  ServerBindingsObjectSolaceBindingVersionValues[ServerBindingsObjectSolaceBindingVersionNumber_0Dot_2Dot_0]: ServerBindingsObjectSolaceBindingVersionNumber_0Dot_2Dot_0,
}
