
package models

// ChannelSchema represents a ChannelSchema model.
type ChannelSchema struct {
  Queue *ChannelSchema
  DeadLetterQueue *ChannelSchema
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion
  AdditionalProperties map[string]interface{}
}