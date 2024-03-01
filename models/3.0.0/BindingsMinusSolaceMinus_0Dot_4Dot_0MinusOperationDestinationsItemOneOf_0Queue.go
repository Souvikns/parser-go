
package models

// BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue represents a BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue model.
type BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue struct {
  Name string `json:"name"`
  TopicSubscriptions []string `json:"topicSubscriptions"`
  AccessType *BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0QueueAccessType `json:"accessType"`
  MaxTtl string `json:"maxTtl"`
  MaxMsgSpoolUsage string `json:"maxMsgSpoolUsage"`
  AdditionalProperties map[string]interface{} `json:"-"`
}