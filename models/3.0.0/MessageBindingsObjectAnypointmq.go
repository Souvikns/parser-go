
package models

// MessageBindingsObjectAnypointmq represents a MessageBindingsObjectAnypointmq model.
type MessageBindingsObjectAnypointmq struct {
  BindingVersion *MessageBindingsObjectAnypointmqBindingVersion `json:"bindingVersion"`
  Headers interface{} `json:"headers"`
  AdditionalProperties map[string]interface{} `json:"-"`
}