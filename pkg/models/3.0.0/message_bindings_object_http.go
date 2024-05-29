
package models

// MessageBindingsObjectHttp represents a MessageBindingsObjectHttp model.
type MessageBindingsObjectHttp struct {
  BindingVersion *MessageBindingsObjectHttpBindingVersion `json:"bindingVersion,omitempty"`
  Headers *Schema `json:"headers,omitempty"`
  StatusCode float64 `json:"statusCode,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}