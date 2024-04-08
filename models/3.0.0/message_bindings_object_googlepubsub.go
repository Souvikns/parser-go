
package models

// MessageBindingsObjectGooglepubsub represents a MessageBindingsObjectGooglepubsub model.
type MessageBindingsObjectGooglepubsub struct {
  BindingVersion *MessageBindingsObjectGooglepubsubBindingVersion `json:"bindingVersion"`
  Attributes map[string]interface{} `json:"attributes"`
  OrderingKey string `json:"orderingKey"`
  Schema *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusMessageSchema `json:"schema"`
  AdditionalProperties map[string]interface{} `json:"-"`
}