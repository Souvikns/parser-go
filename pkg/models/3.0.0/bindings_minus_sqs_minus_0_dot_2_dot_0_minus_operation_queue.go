
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue struct {
  Ref string `json:"$ref"`
  Name string `json:"name"`
  FifoQueue bool `json:"fifoQueue"`
  DeduplicationScope *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope `json:"deduplicationScope"`
  FifoThroughputLimit *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit `json:"fifoThroughputLimit"`
  DeliveryDelay int `json:"deliveryDelay"`
  VisibilityTimeout int `json:"visibilityTimeout"`
  ReceiveMessageWaitTime int `json:"receiveMessageWaitTime"`
  MessageRetentionPeriod int `json:"messageRetentionPeriod"`
  RedrivePolicy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationRedrivePolicy `json:"redrivePolicy"`
  Policy *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationPolicy `json:"policy"`
  Tags map[string]interface{} `json:"tags"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}