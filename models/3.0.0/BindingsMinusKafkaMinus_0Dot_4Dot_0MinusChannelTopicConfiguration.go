
package models

// BindingsMinusKafkaMinus_0Dot_4Dot_0MinusChannelTopicConfiguration represents a BindingsMinusKafkaMinus_0Dot_4Dot_0MinusChannelTopicConfiguration model.
type BindingsMinusKafkaMinus_0Dot_4Dot_0MinusChannelTopicConfiguration struct {
  CleanupDotPolicy []BindingsMinusKafkaMinus_0Dot_4Dot_0MinusChannelTopicConfigurationCleanupDotPolicyItem `json:"cleanup.policy"`
  RetentionDotMs int `json:"retention.ms"`
  RetentionDotBytes int `json:"retention.bytes"`
  DeleteDotRetentionDotMs int `json:"delete.retention.ms"`
  MaxDotMessageDotBytes int `json:"max.message.bytes"`
}