
package models

// ChannelBindingsObjectGooglepubsub represents a ChannelBindingsObjectGooglepubsub model.
type ChannelBindingsObjectGooglepubsub struct {
  BindingVersion *ChannelBindingsObjectGooglepubsubBindingVersion
  Labels map[string]interface{}
  MessageRetentionDuration string
  MessageStoragePolicy *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusChannelMessageStoragePolicy
  SchemaSettings *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusChannelSchemaSettings
  AdditionalProperties map[string]interface{}
}