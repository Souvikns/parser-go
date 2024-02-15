
package models

// MessageBindingsObjectGooglepubsub represents a MessageBindingsObjectGooglepubsub model.
type MessageBindingsObjectGooglepubsub struct {
  BindingVersion *MessageBindingsObjectGooglepubsubBindingVersion
  Attributes map[string]interface{}
  OrderingKey string
  Schema *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusMessageSchema
  AdditionalProperties map[string]interface{}
}