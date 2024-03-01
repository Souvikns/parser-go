
package models

// ChannelBindingsObjectKafka represents a ChannelBindingsObjectKafka model.
type ChannelBindingsObjectKafka struct {
  BindingVersion *ChannelBindingsObjectKafkaBindingVersion `json:"bindingVersion"`
  Topic string `json:"topic"`
  Partitions int `json:"partitions"`
  Replicas int `json:"replicas"`
  TopicConfiguration *BindingsMinusKafkaMinus_0Dot_4Dot_0MinusChannelTopicConfiguration `json:"topicConfiguration"`
  AdditionalProperties map[string]interface{} `json:"-"`
}