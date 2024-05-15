
package models

// MessageBindingsObjectJms represents a MessageBindingsObjectJms model.
type MessageBindingsObjectJms struct {
  BindingVersion *MessageBindingsObjectJmsBindingVersion `json:"bindingVersion"`
  Headers *Schema `json:"headers"`
  AdditionalProperties map[string]interface{} `json:"-"`
}