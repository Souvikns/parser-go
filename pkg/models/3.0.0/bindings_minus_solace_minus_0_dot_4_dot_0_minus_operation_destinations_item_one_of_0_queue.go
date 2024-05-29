
package models

// BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue represents a BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue model.
type BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue struct {
  Name string `json:"name,omitempty"`
  TopicSubscriptions []string `json:"topicSubscriptions,omitempty"`
  AccessType *BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0QueueAccessType `json:"accessType,omitempty"`
  MaxTtl string `json:"maxTtl,omitempty"`
  MaxMsgSpoolUsage string `json:"maxMsgSpoolUsage,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}