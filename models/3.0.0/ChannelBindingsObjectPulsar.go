
package models

// ChannelBindingsObjectPulsar represents a ChannelBindingsObjectPulsar model.
type ChannelBindingsObjectPulsar struct {
  BindingVersion *ChannelBindingsObjectPulsarBindingVersion
  Namespace string
  Persistence *BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelPersistence
  Compaction int
  GeoMinusReplication []string
  Retention *BindingsMinusPulsarMinus_0Dot_1Dot_0MinusChannelRetention
  Ttl int
  Deduplication bool
  AdditionalProperties map[string]interface{}
}