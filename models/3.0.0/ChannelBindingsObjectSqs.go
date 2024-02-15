
package models

// ChannelBindingsObjectSqs represents a ChannelBindingsObjectSqs model.
type ChannelBindingsObjectSqs struct {
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion
  Queue *ChannelSchema
  DeadLetterQueue *ChannelSchema
  AdditionalProperties map[string]interface{}
}