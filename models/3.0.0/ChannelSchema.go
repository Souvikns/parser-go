
package models

// ChannelSchema represents a ChannelSchema model.
type ChannelSchema struct {
  Queue *ChannelSchema `json:"queue"`
  DeadLetterQueue *ChannelSchema `json:"deadLetterQueue"`
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion `json:"bindingVersion"`
  AdditionalProperties map[string]interface{} `json:"-"`
}