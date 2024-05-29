
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue struct {
  Ref string `json:"$ref,omitempty"`
  Name string `json:"name,omitempty"`
  FifoQueue bool `json:"fifoQueue,omitempty"`
  DeduplicationScope *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope `json:"deduplicationScope,omitempty"`
  FifoThroughputLimit *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit `json:"fifoThroughputLimit,omitempty"`
  DeliveryDelay int `json:"deliveryDelay,omitempty"`
  VisibilityTimeout int `json:"visibilityTimeout,omitempty"`
  ReceiveMessageWaitTime int `json:"receiveMessageWaitTime,omitempty"`
  MessageRetentionPeriod int `json:"messageRetentionPeriod,omitempty"`
  RedrivePolicy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationRedrivePolicy `json:"redrivePolicy,omitempty"`
  Policy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationPolicy `json:"policy,omitempty"`
  Tags map[string]interface{} `json:"tags,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}