
package models

// OperationBindingsObjectAmqp represents a OperationBindingsObjectAmqp model.
type OperationBindingsObjectAmqp struct {
  BindingVersion *OperationBindingsObjectAmqpBindingVersion `json:"bindingVersion,omitempty"`
  Expiration int `json:"expiration,omitempty"`
  UserId string `json:"userId,omitempty"`
  Cc []string `json:"cc,omitempty"`
  Priority int `json:"priority,omitempty"`
  DeliveryMode *BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode `json:"deliveryMode,omitempty"`
  Mandatory bool `json:"mandatory,omitempty"`
  Bcc []string `json:"bcc,omitempty"`
  Timestamp bool `json:"timestamp,omitempty"`
  Ack bool `json:"ack,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}