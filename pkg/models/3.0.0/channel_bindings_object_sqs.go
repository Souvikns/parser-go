
package models

// ChannelBindingsObjectSqs represents a ChannelBindingsObjectSqs model.
type ChannelBindingsObjectSqs struct {
  BindingVersion *ChannelBindingsObjectSqsBindingVersion `json:"bindingVersion,omitempty"`
  Queue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue `json:"queue,omitempty"`
  DeadLetterQueue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue `json:"deadLetterQueue,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}