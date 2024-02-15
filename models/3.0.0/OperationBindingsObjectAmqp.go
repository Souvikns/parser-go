
package models

// OperationBindingsObjectAmqp represents a OperationBindingsObjectAmqp model.
type OperationBindingsObjectAmqp struct {
  BindingVersion *OperationBindingsObjectAmqpBindingVersion
  Expiration int
  UserId string
  Cc []string
  Priority int
  DeliveryMode *BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode
  Mandatory bool
  Bcc []string
  Timestamp bool
  Ack bool
  AdditionalProperties map[string]interface{}
}