
package models

// MessageBindingsObjectJms represents a MessageBindingsObjectJms model.
type MessageBindingsObjectJms struct {
  BindingVersion *MessageBindingsObjectJmsBindingVersion `json:"bindingVersion,omitempty"`
  Headers *Schema `json:"headers,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}