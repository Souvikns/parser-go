
package models

// ChannelBindingsObjectAmqpBindingVersion represents an enum of ChannelBindingsObjectAmqpBindingVersion.
type ChannelBindingsObjectAmqpBindingVersion uint

const (
  ChannelBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0 ChannelBindingsObjectAmqpBindingVersion = iota
)

// Value returns the value of the enum.
func (op ChannelBindingsObjectAmqpBindingVersion) Value() any {
	if op >= ChannelBindingsObjectAmqpBindingVersion(len(ChannelBindingsObjectAmqpBindingVersionValues)) {
		return nil
	}
	return ChannelBindingsObjectAmqpBindingVersionValues[op]
}

var ChannelBindingsObjectAmqpBindingVersionValues = []any{"0.3.0"}
var ValuesToChannelBindingsObjectAmqpBindingVersion = map[any]ChannelBindingsObjectAmqpBindingVersion{
  ChannelBindingsObjectAmqpBindingVersionValues[ChannelBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0]: ChannelBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0,
}
