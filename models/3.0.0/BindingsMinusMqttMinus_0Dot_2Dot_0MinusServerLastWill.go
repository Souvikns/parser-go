
package models

// BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill represents a BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill model.
type BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill struct {
  Topic string
  Qos *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos
  Message string
  Retain bool
  AdditionalProperties map[string]interface{}
}