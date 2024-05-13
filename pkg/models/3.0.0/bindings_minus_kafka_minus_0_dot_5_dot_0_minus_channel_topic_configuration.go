
package models

// BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration represents a BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration model.
type BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration struct {
  CleanupDotPolicy []BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfigurationCleanupDotPolicyItem `json:"cleanup.policy"`
  RetentionDotMs int `json:"retention.ms"`
  RetentionDotBytes int `json:"retention.bytes"`
  DeleteDotRetentionDotMs int `json:"delete.retention.ms"`
  MaxDotMessageDotBytes int `json:"max.message.bytes"`
  ConfluentDotKeyDotSchemaDotValidation bool `json:"confluent.key.schema.validation"`
  ConfluentDotKeyDotSubjectDotNameDotStrategy string `json:"confluent.key.subject.name.strategy"`
  ConfluentDotValueDotSchemaDotValidation bool `json:"confluent.value.schema.validation"`
  ConfluentDotValueDotSubjectDotNameDotStrategy string `json:"confluent.value.subject.name.strategy"`
  AdditionalProperties map[string]interface{} `json:"-"`
}