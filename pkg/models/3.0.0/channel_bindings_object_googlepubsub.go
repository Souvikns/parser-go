
package models

// ChannelBindingsObjectGooglepubsub represents a ChannelBindingsObjectGooglepubsub model.
type ChannelBindingsObjectGooglepubsub struct {
  BindingVersion *ChannelBindingsObjectGooglepubsubBindingVersion `json:"bindingVersion,omitempty"`
  Labels map[string]interface{} `json:"labels,omitempty"`
  MessageRetentionDuration string `json:"messageRetentionDuration,omitempty"`
  MessageStoragePolicy *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusChannelMessageStoragePolicy `json:"messageStoragePolicy,omitempty"`
  SchemaSettings *BindingsMinusGooglepubsubMinus_0Dot_2Dot_0MinusChannelSchemaSettings `json:"schemaSettings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}