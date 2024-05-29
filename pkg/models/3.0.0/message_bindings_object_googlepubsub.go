
package models

// MessageBindingsObjectGooglepubsub represents a MessageBindingsObjectGooglepubsub model.
type MessageBindingsObjectGooglepubsub struct {
  BindingVersion *MessageBindingsObjectGooglepubsubBindingVersion `json:"bindingVersion,omitempty"`
  Attributes map[string]interface{} `json:"attributes,omitempty"`
  OrderingKey string `json:"orderingKey,omitempty"`
  Schema *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusMessageSchema `json:"schema,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}