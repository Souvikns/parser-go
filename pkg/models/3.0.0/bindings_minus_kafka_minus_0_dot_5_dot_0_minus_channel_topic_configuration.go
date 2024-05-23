
package models

// BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration represents a BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration model.
type BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration struct {
  CleanupDotPolicy []BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfigurationCleanupDotPolicyItem `json:"cleanup.policy,omitempty"`
  RetentionDotMs int `json:"retention.ms,omitempty"`
  RetentionDotBytes int `json:"retention.bytes,omitempty"`
  DeleteDotRetentionDotMs int `json:"delete.retention.ms,omitempty"`
  MaxDotMessageDotBytes int `json:"max.message.bytes,omitempty"`
  ConfluentDotKeyDotSchemaDotValidation bool `json:"confluent.key.schema.validation,omitempty"`
  ConfluentDotKeyDotSubjectDotNameDotStrategy string `json:"confluent.key.subject.name.strategy,omitempty"`
  ConfluentDotValueDotSchemaDotValidation bool `json:"confluent.value.schema.validation,omitempty"`
  ConfluentDotValueDotSubjectDotNameDotStrategy string `json:"confluent.value.subject.name.strategy,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}