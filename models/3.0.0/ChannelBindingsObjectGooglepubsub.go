
package models

// ChannelBindingsObjectGooglepubsub represents a ChannelBindingsObjectGooglepubsub model.
type ChannelBindingsObjectGooglepubsub struct {
  BindingVersion *ChannelBindingsObjectGooglepubsubBindingVersion `json:"bindingVersion"`
  Labels map[string]interface{} `json:"labels"`
  MessageRetentionDuration string `json:"messageRetentionDuration"`
  MessageStoragePolicy *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusChannelMessageStoragePolicy `json:"messageStoragePolicy"`
  SchemaSettings *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusChannelSchemaSettings `json:"schemaSettings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}