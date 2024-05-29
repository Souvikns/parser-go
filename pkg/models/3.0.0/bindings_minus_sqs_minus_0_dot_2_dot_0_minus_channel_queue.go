
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue struct {
  Name string `json:"name,omitempty"`
  FifoQueue bool `json:"fifoQueue,omitempty"`
  DeduplicationScope *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope `json:"deduplicationScope,omitempty"`
  FifoThroughputLimit *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit `json:"fifoThroughputLimit,omitempty"`
  DeliveryDelay int `json:"deliveryDelay,omitempty"`
  VisibilityTimeout int `json:"visibilityTimeout,omitempty"`
  ReceiveMessageWaitTime int `json:"receiveMessageWaitTime,omitempty"`
  MessageRetentionPeriod int `json:"messageRetentionPeriod,omitempty"`
  RedrivePolicy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy `json:"redrivePolicy,omitempty"`
  Policy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelPolicy `json:"policy,omitempty"`
  Tags map[string]interface{} `json:"tags,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}