
package models

// BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill represents a BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill model.
type BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill struct {
  Topic string `json:"topic"`
  Qos *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos `json:"qos"`
  Message string `json:"message"`
  Retain bool `json:"retain"`
  AdditionalProperties map[string]interface{} `json:"-"`
}