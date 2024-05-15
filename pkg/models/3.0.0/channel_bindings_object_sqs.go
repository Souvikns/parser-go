
package models

// ChannelBindingsObjectSqs represents a ChannelBindingsObjectSqs model.
type ChannelBindingsObjectSqs struct {
  BindingVersion *ChannelBindingsObjectSqsBindingVersion `json:"bindingVersion"`
  Queue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue `json:"queue"`
  DeadLetterQueue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue `json:"deadLetterQueue"`
  AdditionalProperties map[string]interface{} `json:"-"`
}