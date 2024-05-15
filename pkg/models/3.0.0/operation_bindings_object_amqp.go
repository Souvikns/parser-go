
package models

// OperationBindingsObjectAmqp represents a OperationBindingsObjectAmqp model.
type OperationBindingsObjectAmqp struct {
  BindingVersion *OperationBindingsObjectAmqpBindingVersion `json:"bindingVersion"`
  Expiration int `json:"expiration"`
  UserId string `json:"userId"`
  Cc []string `json:"cc"`
  Priority int `json:"priority"`
  DeliveryMode *BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode `json:"deliveryMode"`
  Mandatory bool `json:"mandatory"`
  Bcc []string `json:"bcc"`
  Timestamp bool `json:"timestamp"`
  Ack bool `json:"ack"`
  AdditionalProperties map[string]interface{} `json:"-"`
}