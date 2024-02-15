
package models

// ChannelBindingsObjectWsBindingVersion represents an enum of ChannelBindingsObjectWsBindingVersion.
type ChannelBindingsObjectWsBindingVersion uint

const (
  ChannelBindingsObjectWsBindingVersionNumber_0Dot_1Dot_0 ChannelBindingsObjectWsBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectWsBindingVersion) Value() any {
	if op >= ChannelBindingsObjectWsBindingVersion(len(ChannelBindingsObjectWsBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectWsBindingVersionValues[op]
}

var ChannelBindingsObjectWsBindingVersionValues = []any{"0.1.0"}
var ValuesToChannelBindingsObjectWsBindingVersion = map[any]ChannelBindingsObjectWsBindingVersion{
  ChannelBindingsObjectWsBindingVersionValues[ChannelBindingsObjectWsBindingVersionNumber_0Dot_1Dot_0]: ChannelBindingsObjectWsBindingVersionNumber_0Dot_1Dot_0,
}
