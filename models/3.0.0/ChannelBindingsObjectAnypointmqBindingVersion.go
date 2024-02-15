
package models

// ChannelBindingsObjectAnypointmqBindingVersion represents an enum of ChannelBindingsObjectAnypointmqBindingVersion.
type ChannelBindingsObjectAnypointmqBindingVersion uint

const (
  ChannelBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1 ChannelBindingsObjectAnypointmqBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectAnypointmqBindingVersion) Value() any {
	if op >= ChannelBindingsObjectAnypointmqBindingVersion(len(ChannelBindingsObjectAnypointmqBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectAnypointmqBindingVersionValues[op]
}

var ChannelBindingsObjectAnypointmqBindingVersionValues = []any{"0.0.1"}
var ValuesToChannelBindingsObjectAnypointmqBindingVersion = map[any]ChannelBindingsObjectAnypointmqBindingVersion{
  ChannelBindingsObjectAnypointmqBindingVersionValues[ChannelBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1]: ChannelBindingsObjectAnypointmqBindingVersionNumber_0Dot_0Dot_1,
}
