
package models

// MessageBindingsObjectJms represents a MessageBindingsObjectJms model.
type MessageBindingsObjectJms struct {
  BindingVersion *MessageBindingsObjectJmsBindingVersion `json:"bindingVersion"`
  Headers interface{} `json:"headers"`
  AdditionalProperties map[string]interface{} `json:"-"`
}