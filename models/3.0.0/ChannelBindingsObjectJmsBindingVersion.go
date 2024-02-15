
package models

// ChannelBindingsObjectJmsBindingVersion represents an enum of ChannelBindingsObjectJmsBindingVersion.
type ChannelBindingsObjectJmsBindingVersion uint

const (
  ChannelBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1 ChannelBindingsObjectJmsBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectJmsBindingVersion) Value() any {
	if op >= ChannelBindingsObjectJmsBindingVersion(len(ChannelBindingsObjectJmsBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectJmsBindingVersionValues[op]
}

var ChannelBindingsObjectJmsBindingVersionValues = []any{"0.0.1"}
var ValuesToChannelBindingsObjectJmsBindingVersion = map[any]ChannelBindingsObjectJmsBindingVersion{
  ChannelBindingsObjectJmsBindingVersionValues[ChannelBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1]: ChannelBindingsObjectJmsBindingVersionNumber_0Dot_0Dot_1,
}
