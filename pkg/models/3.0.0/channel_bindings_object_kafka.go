
package models

// ChannelBindingsObjectKafka represents a ChannelBindingsObjectKafka model.
type ChannelBindingsObjectKafka struct {
  BindingVersion *ChannelBindingsObjectKafkaBindingVersion `json:"bindingVersion,omitempty"`
  Topic string `json:"topic,omitempty"`
  Partitions int `json:"partitions,omitempty"`
  Replicas int `json:"replicas,omitempty"`
  TopicConfiguration *BindingsMinusKafkaMinus_0Dot_5Dot_0MinusChannelTopicConfiguration `json:"topicConfiguration,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}