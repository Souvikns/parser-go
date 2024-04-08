
package models

// ChannelBindingsObjectSqs represents a ChannelBindingsObjectSqs model.
type ChannelBindingsObjectSqs struct {
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion `json:"bindingVersion"`
  Queue *ChannelSchema `json:"queue"`
  DeadLetterQueue *ChannelSchema `json:"deadLetterQueue"`
  AdditionalProperties map[string]interface{} `json:"-"`
}