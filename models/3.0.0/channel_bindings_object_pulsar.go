
package models

// ChannelBindingsObjectPulsar represents a ChannelBindingsObjectPulsar model.
type ChannelBindingsObjectPulsar struct {
  BindingVersion *ChannelBindingsObjectPulsarBindingVersion `json:"bindingVersion"`
  Namespace string `json:"namespace"`
  Persistence *BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence `json:"persistence"`
  Compaction int `json:"compaction"`
  GeoMinusReplication []string `json:"geo-replication"`
  Retention *BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelRetention `json:"retention"`
  Ttl int `json:"ttl"`
  Deduplication bool `json:"deduplication"`
  AdditionalProperties map[string]interface{} `json:"-"`
}