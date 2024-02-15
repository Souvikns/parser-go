
package models

// ChannelBindingsObjectKafka represents a ChannelBindingsObjectKafka model.
type ChannelBindingsObjectKafka struct {
  BindingVersion *ChannelBindingsObjectKafkaBindingVersion
  Topic string
  Partitions int
  Replicas int
  TopicConfiguration *BindingsMinusKafkaMinus_0Dot_4Dot_0MinusChannelTopicConfiguration
  AdditionalProperties map[string]interface{}
}