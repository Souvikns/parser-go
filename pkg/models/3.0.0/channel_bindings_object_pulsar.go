
package models

// ChannelBindingsObjectPulsar represents a ChannelBindingsObjectPulsar model.
type ChannelBindingsObjectPulsar struct {
  BindingVersion *ChannelBindingsObjectPulsarBindingVersion `json:"bindingVersion,omitempty"`
  Namespace string `json:"namespace,omitempty"`
  Persistence *BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence `json:"persistence,omitempty"`
  Compaction int `json:"compaction,omitempty"`
  GeoMinusReplication []string `json:"geo-replication,omitempty"`
  Retention *BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelRetention `json:"retention,omitempty"`
  Ttl int `json:"ttl,omitempty"`
  Deduplication bool `json:"deduplication,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}