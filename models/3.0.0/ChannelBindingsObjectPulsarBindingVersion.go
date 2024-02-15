
package models

// ChannelBindingsObjectPulsarBindingVersion represents an enum of ChannelBindingsObjectPulsarBindingVersion.
type ChannelBindingsObjectPulsarBindingVersion uint

const (
  ChannelBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0 ChannelBindingsObjectPulsarBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectPulsarBindingVersion) Value() any {
	if op >= ChannelBindingsObjectPulsarBindingVersion(len(ChannelBindingsObjectPulsarBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectPulsarBindingVersionValues[op]
}

var ChannelBindingsObjectPulsarBindingVersionValues = []any{"0.1.0"}
var ValuesToChannelBindingsObjectPulsarBindingVersion = map[any]ChannelBindingsObjectPulsarBindingVersion{
  ChannelBindingsObjectPulsarBindingVersionValues[ChannelBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0]: ChannelBindingsObjectPulsarBindingVersionNumber_0Dot_1Dot_0,
}
