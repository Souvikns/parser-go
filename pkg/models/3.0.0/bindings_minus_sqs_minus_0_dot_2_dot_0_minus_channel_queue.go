
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueue struct {
  Name string `json:"name"`
  FifoQueue bool `json:"fifoQueue"`
  DeduplicationScope *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope `json:"deduplicationScope"`
  FifoThroughputLimit *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit `json:"fifoThroughputLimit"`
  DeliveryDelay int `json:"deliveryDelay"`
  VisibilityTimeout int `json:"visibilityTimeout"`
  ReceiveMessageWaitTime int `json:"receiveMessageWaitTime"`
  MessageRetentionPeriod int `json:"messageRetentionPeriod"`
  RedrivePolicy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy `json:"redrivePolicy"`
  Policy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelPolicy `json:"policy"`
  Tags map[string]interface{} `json:"tags"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}