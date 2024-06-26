
package models

// MessageBindingsObjectHttp represents a MessageBindingsObjectHttp model.
type MessageBindingsObjectHttp struct {
  BindingVersion *MessageBindingsObjectHttpBindingVersion `json:"bindingVersion"`
  Headers *Schema `json:"headers"`
  StatusCode float64 `json:"statusCode"`
  AdditionalProperties map[string]interface{} `json:"-"`
}