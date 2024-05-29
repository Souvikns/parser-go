
package models

// BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill represents a BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill model.
type BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill struct {
  Topic string `json:"topic,omitempty"`
  Qos *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos `json:"qos,omitempty"`
  Message string `json:"message,omitempty"`
  Retain bool `json:"retain,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}